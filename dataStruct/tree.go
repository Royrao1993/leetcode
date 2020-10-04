package dataStruct

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
	}
}
