package main

import (
	"./dataStruct"
	"./linkedlist"
	"fmt"
)

func main() {
	head := &dataStruct.ListNode{Val: 1}
	res := linkedlist.RecordList(head)
	fmt.Println(res)
}
