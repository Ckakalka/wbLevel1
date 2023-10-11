package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

var signals chan os.Signal

func init() {
	signals = make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
}

func checkArgs() bool {
	if len(os.Args) != 2 {
		return false
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return false
	}
	if !(n >= 1) {
		return false
	}
	return true
}

func StartWorker(dataChan chan int, stdOutChan chan struct{}) {
	for data := range dataChan {
		<-stdOutChan
		fmt.Println(data)
		stdOutChan <- struct{}{}
	}
}

func main() {
	if !checkArgs() {
		fmt.Printf("Invalid input arguments\nUsage:\nprogram N\nwhere N - natural number.\n")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])

	stdOutChan := make(chan struct{}, 1)
	stdOutChan <- struct{}{}
	dataChan := make(chan int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			StartWorker(dataChan, stdOutChan)
		}()
	}
	quit := make(chan struct{})
	go func() {
		<-signals
		<-stdOutChan
		fmt.Println("SigInt processed...")
		stdOutChan <- struct{}{}
		quit <- struct{}{}
	}()
loop:
	for {
		select {
		case <-quit:
			break loop
		default:
			dataChan <- 123456789
		}
	}
	<-stdOutChan
	fmt.Println("Completing the goroutines...")
	stdOutChan <- struct{}{}
	close(dataChan)
	wg.Wait()
}
