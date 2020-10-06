package main

import (
	"./design"
	"fmt"
)

func main() {
	s := design.NewRandomizedSet()
	s.Insert(0)
	s.Insert(1)
	s.Remove(0)
	s.Insert(2)
	s.Remove(1)
	fmt.Println(s.GetRandom())
}
