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
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	prev := root
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head != nil {
			if prev == nil {
				return false
			}
			queue = append(queue, head.Left)
			queue = append(queue, head.Right)
		}

		prev = head

	}
	return true
}

// @leet end

