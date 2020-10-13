package linkedlist

import "../dataStruct"

func SwapPairs(head *dataStruct.ListNode) *dataStruct.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &dataStruct.ListNode{Val: -1}
	cur := head
	dummy.Next = cur
	prev := dummy
	for cur != nil && cur.Next != nil {
		next := cur.Next
		prev.Next = next
		cur.Next = next.Next
		next.Next = cur
		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}
