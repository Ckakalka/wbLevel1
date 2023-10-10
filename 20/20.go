package main

import (
	"fmt"
	"strings"
)

func ReverseWordsInStr(s string) string {
	var result strings.Builder
	arrWords := strings.Split(s, " ")
	if len(arrWords) == 0 {
		return ""
	}
	for i := len(arrWords) - 1; i > 0; i-- {
		result.WriteString(arrWords[i])
		result.WriteString(" ")
	}
	result.WriteString(arrWords[0])
	return result.String()
}

func main() {
	str := "snow dog sun"
	rStr := ReverseWordsInStr(str)
	fmt.Println(rStr)
}
