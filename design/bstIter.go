package design

import (
	"../dataStruct"
)

type BSTIterator struct {
	stack []int
}

func NewBSTIter(root *dataStruct.TreeNode) BSTIterator {
	i := BSTIterator{
		stack: []int{},
	}
	inorder(root, &i.stack)
	return i
}

func inorder(root *dataStruct.TreeNode, s *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, s)
	}
	*s = append(*s, root.Val)
	if root.Right != nil {
		inorder(root.Right, s)
	}
}

func (i *BSTIterator) Next() int {
	res := i.stack[0]
	i.stack = i.stack[1:]
	return res
}

func (i *BSTIterator) HasNext() bool {
	return len(i.stack) > 0
}
