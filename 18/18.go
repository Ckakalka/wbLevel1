package main

import (
	"fmt"
	"sync"
)

type ConcurrentCounter struct {
	mx    sync.RWMutex
	value int
}

func (c *ConcurrentCounter) GetValue() int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.value
}

func (c *ConcurrentCounter) Inc() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
}

func main() {
	var wg sync.WaitGroup
	var c ConcurrentCounter
	const countWorker = 10
	const countIterations = 1_000_000
	for i := 0; i < countWorker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < countIterations; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	value := c.GetValue()
	exptectedValue := countWorker * countIterations
	if value != exptectedValue {
		fmt.Printf("error: counter.value=%d, exptected %d\n", value, exptectedValue)
		return
	}
	fmt.Println(c.GetValue())
}
