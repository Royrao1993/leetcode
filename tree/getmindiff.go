package tree

import (
	"../dataStruct"
	"../utils"
	"math"
)

var (
	res  int
	prev int
)

func getMinDiff(root *dataStruct.TreeNode) int {
	res = math.MaxInt64
	prev = -1
	help(root)
	return res
}

func help(node *dataStruct.TreeNode) {
	if node == nil {
		return
	}
	help(node.Left)
	if prev >= 0 {
		res = utils.Min(res, node.Val-prev)
	}
	prev = node.Val
	help(node.Right)
}
