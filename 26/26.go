package main

import (
	"fmt"
	"unicode"
)

type runeSet map[rune]struct{}

func isUniqueSymbs(s string) bool {
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
	if isUniqueSymbs(str) {
		fmt.Printf("error: %s\n", str)
	}
	str = "АБСДЕЙЦУFR"
	if !isUniqueSymbs(str) {
		fmt.Printf("error: %s\n", str)
	}
}
