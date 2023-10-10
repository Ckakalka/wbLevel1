package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mx sync.RWMutex
	m  map[int]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[int]int),
	}
}

func (cMap *ConcurrentMap) Load(key int) (int, bool) {
	cMap.mx.RLock()
	defer cMap.mx.RUnlock()
	data, ok := cMap.m[key]
	return data, ok
}

func (cMap *ConcurrentMap) Store(key int, data int) {
	cMap.mx.Lock()
	defer cMap.mx.Unlock()
	cMap.m[key] = data
}

func (cMap *ConcurrentMap) Inc(key int, incValue int) {
	cMap.mx.Lock()
	defer cMap.mx.Unlock()
	value := cMap.m[key]
	cMap.m[key] = value + incValue
}

func main() {
	var wg sync.WaitGroup
	cMap := NewConcurrentMap()
	startDataWriter := func(start, end, step int) {
		for i := start; i < end; i += step {
			cMap.Inc(i, 1)
		}
	}
	left := 0
	right := 1_000_000
	const countWriter = 5
	for i := 0; i < countWriter; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			startDataWriter(left, right, 1)
		}()

	}
	wg.Wait()
	if len(cMap.m) != right-left {
		fmt.Printf("error: len(cMap.m)=%d, expected %d\n", len(cMap.m), right-left)
	}
	for k, v := range cMap.m {
		if v != countWriter {
			fmt.Printf("error: key=%d value=%d (expected value=2)\n", k, v)
		}
	}
}
