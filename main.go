package main

import (
	"./dataStruct"
	"./design"
	"fmt"
)

func main() {
	left := dataStruct.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	right := dataStruct.TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	root := dataStruct.TreeNode{
		Val:   7,
		Left:  &left,
		Right: &right,
	}
	iter := design.NewBSTIter(&root)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
}
