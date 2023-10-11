package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	squareChan := make(chan int, 2)
	SquareToChan := func(n int) {
		squareChan <- n * n
	}
	for i := range arr {
		go SquareToChan(arr[i])
	}
	for i := 0; i < len(arr); i++ {
		square := <-squareChan
		fmt.Printf("%d ", square)
	}
	fmt.Println()
}
