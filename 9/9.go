package main

import "fmt"

func inputWriter(s []int, inputChan chan<- int) {
	for _, v := range s {
		inputChan <- v
	}
	close(inputChan)
}

func inputReader(inputChan <-chan int, outputChan chan<- int) {
	for v := range inputChan {
		outputChan <- 2 * v
	}
	close(outputChan)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	inputChan := make(chan int)
	go inputWriter(arr[:], inputChan)
	outputChan := make(chan int)
	go inputReader(inputChan, outputChan)
	for v := range outputChan {
		fmt.Println(v)
	}
}
