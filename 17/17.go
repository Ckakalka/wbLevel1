package main

import "fmt"

// Time complexity: O(log(n))
// Memory complexity: O(1)
func BinSearch(arr []int, key int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		midValue := arr[mid]
		if midValue < key {
			left = mid + 1
		} else if midValue > key {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range data {
		if !BinSearch(data, v) {
			fmt.Printf("error value=%d not found\n", v)
		}
	}
	if BinSearch(data, 15) {
		fmt.Println("error value=15 found")
	}
}
