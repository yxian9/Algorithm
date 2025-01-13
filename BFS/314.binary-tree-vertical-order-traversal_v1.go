package leetcode

import (
	"sort"
)

// @leet start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type item struct {
	val, col, row int
}
type s314 struct {
	order      []item
	minC, maxC int
}

func (s *s314) dfs(node *TreeNode, r, c int) {
	if node == nil {
		return
	}
	s.maxC = max(s.maxC, c)
	s.minC = min(s.minC, c)
	s.order = append(s.order, item{
		val: node.Val,
		col: c,
		row: r,
	})
	s.dfs(node.Left, r+1, c-1)
	s.dfs(node.Right, r+1, c+1)
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	s := s314{}
	s.dfs(root, 0, 0)

	var res [][]int
	sort.SliceStable(s.order, func(a, b int) bool {
		return s.order[a].col < s.order[b].col
	})
	idx := 0
	for c := s.minC; c <= s.maxC; c++ {
		var curCol []item
		for ; idx < len(s.order); idx++ {
			if s.order[idx].col != c {
				break
			}
			curCol = append(curCol, s.order[idx])
		}
		sort.SliceStable(curCol, func(a, b int) bool {
			return curCol[a].row < curCol[b].row
		})
		var temp []int
		for _, v := range curCol {
			temp = append(temp, v.val)
		}
		res = append(res, temp)
	}

	return res
}

// @leet end
