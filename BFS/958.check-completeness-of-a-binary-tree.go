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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	prev := root
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head != nil { // check nil here
			if prev == nil {
				return false
			}
			queue = append(queue, head.Left) // specifically push nil into queue
			queue = append(queue, head.Right)
		}

		prev = head

	}
	return true
}

// @leet end

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
		if head == nil { // this also work. but hard to understand
			prev = head
			continue
		}

		if prev == nil {
			return false
		}
		queue = append(queue, head.Left)
		queue = append(queue, head.Right)

	}
	return true
}

func isCompleteTree(root *TreeNode) bool {
	count, id := -1, 0 // NOTE: 1, 0, count need -1, id start from 0 is fixed
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, n int) {
		if node == nil {
			return
		}
		id = max(id, n)
		count++
		dfs(node.Left, 2*n+1)
		dfs(node.Right, 2*n+2)
	}
	dfs(root, 0)
	return count == id
}

// @leet end

func isCompleteTree2(root *TreeNode) bool { // NOTE: logic flaw. dfs is hard to aware the status on other branch
	var (
		first = true
		dfs   func(*TreeNode) bool
	)
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Left == nil && node.Right == nil {
			return true
		}
		if node.Left != nil && node.Right != nil {
			return dfs(node.Left) && dfs(node.Right)
		}
		if node.Left == nil {
			return false
		}
		if first {
			first = false
			return dfs(node.Left)
		} else {
			return false
		}
	}

	return dfs(root)
}


