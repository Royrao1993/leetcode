package main

import (
	"./dataStruct"
	"./linkedlist"
	"fmt"
)

func main() {
	head := &dataStruct.ListNode{Val: 1}
	node1 := &dataStruct.ListNode{Val: 2}
	node2 := &dataStruct.ListNode{Val: 3}
	node3 := &dataStruct.ListNode{Val: 4}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	h := linkedlist.SwapPairs(head)
	fmt.Println(h.Val)
}
