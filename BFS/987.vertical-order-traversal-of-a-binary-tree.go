package leetcode

import "sort"

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type node1 struct {
	*TreeNode
	col, row int
}
type node2 struct {
	val, row int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		l, r int
		ret  [][]int
	)
	mp := map[int][]node2{}
	queue := []node1{{root, 0, 0}}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		mp[top.col] = append(mp[top.col], node2{top.Val, top.row})

		if top.Left != nil {
			l = min(l, top.col-1)
			queue = append(queue, node1{top.Left, top.col - 1, top.row + 1})
		}
		if top.Right != nil {
			r = max(r, top.col+1)
			queue = append(queue, node1{top.Right, top.col + 1, top.row + 1})
		}
	}
	for i := l; i <= r; i++ {
		cur := mp[i]
		// sort cur
		sort.Slice(cur, func(i, j int) bool {
			l, r := cur[i], cur[j]
			if l.row != r.row {
				return l.row < r.row
			}
			return l.val < r.val
		})
		temp := make([]int, len(cur))
		for i, v := range cur {
			temp[i] = v.val
		}
		ret = append(ret, temp)
	}

	return ret
}

// @leet end

