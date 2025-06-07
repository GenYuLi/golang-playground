package main

import (
	"data/util"
	"fmt"
)

func main() {
	s := make([]int, 20)
	for idx, val := range s {
		s[idx] = idx*2 + val
	}
	s2 := s[:]
	s2[0] = 100
	fmt.Printf("s: %v, s2: %v\n", s, s2)

	p := util.Dog{}
	_ = p.Speak() // Call the Speak method of Dog
	println(util.Add(20, 30.0))
}
