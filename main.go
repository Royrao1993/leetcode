package main

import (
	"./dataStruct"
	"./linkedlist"
	"fmt"
)

func main() {
	head := &dataStruct.ListNode{Val: 1}
	tail := &dataStruct.ListNode{Val: 2}
	head.Next = tail
	tail.Next = head
	node := linkedlist.DetectCycle(head)
	fmt.Println(node)
}
