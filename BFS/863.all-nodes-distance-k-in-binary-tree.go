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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	childParMap := map[*TreeNode]*TreeNode{}

	// DFS annotates the parent
	var dfs func(child, par *TreeNode)
	dfs = func(child, par *TreeNode) {
		if child == nil {
			return
		}
		childParMap[child] = par
		dfs(child.Left, child)
		dfs(child.Right, child)
	}

	dfs(root, nil)

	q := []*TreeNode{target}
	seen := map[int]bool{}
	seen[target.Val] = true
	var dist int
	var res []int = []int{}

	// BFS find the nodes at distance k from target node.
	for len(q) > 0 {
		if dist == K {
			for _, v := range q {
				res = append(res, v.Val)
			}
			return res
		}
		dist++
		length := len(q)
		for range length {
			head := q[0]
			q = q[1:]
			// if seen[head.Val] {
			// 	continue
			// }
			if head.Left != nil && !seen[head.Left.Val] {
				// all bsf seen happend during append step !!
				seen[head.Left.Val] = true
				q = append(q, head.Left)
			}
			if head.Right != nil && !seen[head.Right.Val] {
				seen[head.Right.Val] = true
				q = append(q, head.Right)
			}
			parNode := childParMap[head]
			if parNode != nil && !seen[parNode.Val] {
				seen[parNode.Val] = true
				q = append(q, parNode)
			}
		}
	}
	return []int{}
}

// @leet end

