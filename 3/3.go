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

func GetArrPartition(s []int) (countSubslices int, countElemsSubslice int) {
	countSubslices = int(math.Log2(float64(len(s))))
	if countSubslices == 0 {
		countSubslices = 1
	}
	countElemsSubslice = len(s) / countSubslices
	if len(s)%countSubslices != 0 {
		countElemsSubslice++
	}
	return countSubslices, countElemsSubslice
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	countSubslices, countElemsSubslice := GetArrPartition(arr)
	squareChan := make(chan int, 2)
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
