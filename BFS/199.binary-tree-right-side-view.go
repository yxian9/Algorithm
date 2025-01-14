package leetcode

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		lastItem := queue[length-1]
		res = append(res, lastItem.Val)
		for length > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			length--
		}
	}
	return res
}

// @leet end

