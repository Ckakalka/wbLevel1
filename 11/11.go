package main

import "fmt"

type SetStr map[string]struct{}

func (s1 SetStr) Intersection(s2 SetStr) SetStr {
	result := make(SetStr)
	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			result[k] = struct{}{}
		}
	}
	return result
}

func main() {
	s1 := SetStr{"cat": struct{}{}, "dog": struct{}{}, "tirex": struct{}{}}
	s2 := SetStr{"rabbit": struct{}{}, "tirex": struct{}{}, "pikachu": struct{}{}}
	s := s1.Intersection(s2)
	fmt.Println(s)
}
