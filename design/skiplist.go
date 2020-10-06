package design

import (
	"math"
	"math/rand"
)

const (
	maxLevel = 16
	maxRand  = 65535.0
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

type SkipNode struct {
	Val   int
	Right *SkipNode
	Down  *SkipNode
}

type SkipList struct {
	Head *SkipNode
}

func NewSkipList() *SkipList {
	left := make([]*SkipNode, maxLevel)
	right := make([]*SkipNode, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &SkipNode{
			Val:   -1,
			Right: nil,
			Down:  nil,
		}
		right[i] = &SkipNode{
			Val:   20001, // The limit of value in this question
			Right: nil,
			Down:  nil,
		}
	}
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].Right = right[i]
		left[i].Down = left[i+1]
		right[i].Down = right[i+1]
	}
	left[maxLevel-1].Right = right[maxLevel-1]
	return &SkipList{
		Head: left[0],
	}
}

func (l *SkipList) Search(target int) bool {
	node := l.Head
	for node != nil {
		if node.Right.Val > target {
			node = node.Down
		} else if node.Right.Val < target {
			node = node.Right
		} else {
			return true
		}
	}
	return false
}

func (l *SkipList) Add(num int) {
	prev := make([]*SkipNode, maxLevel)
	i := 0
	node := l.Head
	// find the position to insert the value
	for node != nil {
		if node.Right.Val >= num {
			prev[i] = node
			i++
			node = node.Down
		} else {
			node = node.Right
		}
	}
	level := randLevel()
	arr := make([]*SkipNode, level)
	tmp := &SkipNode{-1, nil, nil}
	for i, v := range arr {
		p := prev[maxLevel-level+i]
		v = &SkipNode{num, p.Right, nil}
		p.Right = v
		tmp.Down = v
		tmp = v
	}
}

func (l *SkipList) Erase(num int) bool {
	flag := false
	node := l.Head
	for node != nil {
		if node.Right.Val > num {
			node = node.Down
		} else if node.Right.Val < num {
			node = node.Right
		} else {
			// find the target node to erase, the target is the Right node of current node
			flag = true
			node.Right = node.Right.Right
			node = node.Down
		}
	}
	return flag
}
