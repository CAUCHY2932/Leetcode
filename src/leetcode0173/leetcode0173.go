package leetcode0173

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
type BSTIterator struct {
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{}
}

func (this *BSTIterator) Next() int {
	return 0
}

func (this *BSTIterator) HasNext() bool {
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func next(root *TreeNode) {

	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left != nil {

		return
	}

}

var nums []int

func inOrder(root *TreeNode) {

	if root == nil {
		return
	}
	inOrder(root.Left)
	nums = append(nums, root.Val)
	inOrder(root.Right)

}
