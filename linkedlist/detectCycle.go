package linkedlist

import (
	"../dataStruct"
)

func DetectCycle(head *dataStruct.ListNode) *dataStruct.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}
