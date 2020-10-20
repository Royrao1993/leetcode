package linkedlist

import (
	"../dataStruct"
)

func RecordList(head *dataStruct.ListNode) *dataStruct.ListNode {
	dummy := &dataStruct.ListNode{
		Val:  -1,
		Next: nil,
	}
	dummy.Next = head
	fast := head
	slow := head
	prev := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	// split the linkedlist into two parts
	prev.Next = nil
	prev = nil
	// reverse the second part of linkedlist
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	cur := dummy
	if head == prev {
		return dummy.Next
	}
	for head != nil && prev != nil {
		cur.Next = head
		head = head.Next
		cur = cur.Next
		cur.Next = prev
		prev = prev.Next
		cur = cur.Next
	}
	if prev != nil {
		cur.Next = prev
	}
	return dummy.Next
}
