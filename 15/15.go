package main

import (
	"unsafe"
)

var justString string

func createHugeString(n int) string {
	p := make([]byte, n)
	s := *(*string)(unsafe.Pointer(&p))
	return s
}

func someFunc() {
	str := createHugeString(1 << 10)
	justStringLen := 100
	if len(str) >= justStringLen {
		justString = str[:justStringLen]
	}
}

func main() {
	someFunc()
}
