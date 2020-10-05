package main

import (
	"./design"
	"fmt"
)

func main() {
	t := design.NewTrie()
	t.Insert("hello")
	fmt.Println(t.Search("hello"))
	fmt.Println(t.StartsWith("he"))
	fmt.Println(t.StartsWith("llo"))
}
