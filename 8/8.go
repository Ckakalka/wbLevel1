package main

import (
	"fmt"
)

func setBitI64(a *int64, n int, isSet bool) {
	if !(0 <= n && n < 64) {
		return
	}
	if isSet {
		*a |= (1 << n)
	} else {
		*a &= ^(1 << n)
	}
}

func main() {
	var a int64 = 122
	fmt.Println(a)
	setBitI64(&a, 0, true)
	fmt.Println(a)
}
