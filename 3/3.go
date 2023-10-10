package main

import (
	"fmt"
	"math"
)

func SquareSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v * v
	}
	c <- sum
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	countSubslices := int(math.Log2(float64(len(arr))))
	if countSubslices == 0 {
		countSubslices = 1
	}
	squareChan := make(chan int)
	countElemsSubslice := len(arr) / countSubslices
	if len(arr)%countSubslices != 0 {
		countElemsSubslice++
	}
	for i := 0; i < countSubslices; i++ {
		left := i * countElemsSubslice
		right := (i + 1) * countElemsSubslice
		if right > len(arr) {
			right = len(arr)
		}
		go SquareSum(arr[left:right], squareChan)

	}
	sum := 0
	for i := 0; i < countSubslices; i++ {
		sum += <-squareChan
	}
	fmt.Println(sum)
}
