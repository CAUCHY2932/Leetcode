package main

import (
	"fmt"
	"sort"
)

func main() {
	// c := make(chan int)

	// c <- 6

	// fmt.Println(<-c)

	b := BuildTree([]int{3, 9, 20, -100, -100, 15, 7}, 0)
	fmt.Printf("[res] get buildTee %#v\n", b)
	InOrderPrint(b)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res = make([][]int, 0)
	que := newQueue()
	que.push(root)
	turnLeft := 1 // false
	for !que.isEmpty() {
		curRes := make([]int, 0)
		curSize := len(*que)
		for i := 0; i < curSize; i++ {
			item := que.pop()
			curRes = append(curRes, item.Val)
			if turnLeft == 1 {
				insertNotNil(item.Left, que)
				insertNotNil(item.Right, que)
			} else {
				insertNotNil(item.Right, que)
				insertNotNil(item.Left, que)
			}
		}
		turnLeft = 1 - turnLeft // change
		res = append(res, curRes)
	}
	return res

}

var res [][]int

func insertNotNil(root *TreeNode, que *queue) {
	if root != nil {
		que.push(root)
		fmt.Printf("[proc] insert item %#v", root)
	}
}

type queue []*TreeNode

func (q *queue) push(item *TreeNode) {
	*q = append(*q, item)
}

func (q *queue) pop() *TreeNode {
	// item := (*q)[0]
	item := (*q)[len(*q)-1]
	// *q = (*q)[1:]
	*q = (*q)[:len(*q)-1]
	return item
}
func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func newQueue() *queue {
	return &queue{}
}
