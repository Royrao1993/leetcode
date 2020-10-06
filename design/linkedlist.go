package design

type MyLinkedList struct {
	Len  int
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func NewLinkedList() *MyLinkedList {
	l := &MyLinkedList{
		Len:  0,
		Head: &ListNode{0, nil, nil},
		Tail: &ListNode{0, nil, nil},
	}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
	return l
}

func (l *MyLinkedList) AddHead(val int) {
	l.AddAt(0, val)
}

func (l *MyLinkedList) AddTail(val int) {
	l.AddAt(l.Len, val)
}

func (l *MyLinkedList) AddAt(index, val int) {
	if index < 0 {
		index = 0
	}
	if index > l.Len {
		return
	}
	if index == l.Len {
		// insert at last
		node := &ListNode{val, l.Tail, l.Tail.Prev}
		l.Tail.Prev.Next = node
		l.Tail.Prev = node
	} else {
		next := l.getNode(index)
		prev := next.Prev
		node := &ListNode{val, next, prev}
		prev.Next = node
		next.Prev = node
	}
	l.Len++
}

func (l *MyLinkedList) GetAt(index int) int {
	node := l.getNode(index)
	if node != nil {
		return node.Val
	}
	return -1
}

func (l *MyLinkedList) getNode(index int) *ListNode {
	if index < 0 || index >= l.Len {
		return nil
	}
	node := l.Head
	for index >= 0 {
		node = node.Next
		index--
	}
	return node
}

func (l *MyLinkedList) DeleteAt(index int) {
	node := l.getNode(index)
	if node != nil {
		next := node.Next
		prev := node.Prev
		prev.Next = next
		next.Prev = prev
		l.Len--
	}
}
