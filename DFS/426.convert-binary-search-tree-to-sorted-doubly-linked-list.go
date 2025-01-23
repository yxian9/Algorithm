package leetcode

// @leet start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}
	var head, tail *Node
	var dfs func(*Node)
	dfs = func(cur *Node) {
		if cur == nil {
			return
		}
		dfs(cur.Left)
		if head == nil {
			head = cur
		}
		if tail != nil {
			cur.Left = tail // order does not matter
			tail.Right = cur
		}
		tail = cur

		dfs(cur.Right)
	}
	dfs(root)
	head.Left = tail
	tail.Right = head
	return head
}

// @leet end

