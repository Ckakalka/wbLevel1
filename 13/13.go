package main

import "fmt"

func main() {
	a := 5
	b := -10
	fmt.Printf("a=%d, b=%d\n", a, b)

	// 1 way
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)

	// 2 way
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%d, b=%d\n", a, b)

	// 3 way
	a = a - b
	b = a + b
	a = b - a
	fmt.Printf("a=%d, b=%d\n", a, b)

	// 4 way
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a=%d, b=%d\n", a, b)
}
