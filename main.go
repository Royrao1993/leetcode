package main

import (
	"./string"
	"fmt"
)

func main() {
	res := string.BackspaceCompare("ab#c", "ad#c")
	fmt.Println(res)
}
