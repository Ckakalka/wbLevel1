package main

import "fmt"

func Remove(s []int, i int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:i]...)
	ret = append(ret, s[i+1:]...)
	return ret
}

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d0 := Remove(data[:], 1)
	fmt.Printf("d0=  %v\n", d0)
	fmt.Printf("data=%v\n", data)
	d3 := Remove(data[:], 3)
	fmt.Printf("d3=  %v\n", d3)
	fmt.Printf("data=%v\n", data)
}
