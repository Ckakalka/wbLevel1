package main

import (
	"fmt"
	"unicode"
)

type runeSet map[rune]struct{}

func isSymbsUnique(s string) bool {
	runes := make(runeSet)
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, ok := runes[r]; ok {
			return false
		} else {
			runes[r] = struct{}{}
		}
	}
	return true
}

func main() {
	str := "АБСДа"
	if isSymbsUnique(str) {
		fmt.Printf("error: %s\n", str)
	}
	str = "АБСДЕЙЦУFR"
	if !isSymbsUnique(str) {
		fmt.Printf("error: %s\n", str)
	}
}
