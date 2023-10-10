package main

import "fmt"

func main() {
	a := 5
	b := -10
	fmt.Printf("a=%d, b=%d\n", a, b)

	//  a, b = b, a
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a=%d, b=%d\n", a, b)
}
