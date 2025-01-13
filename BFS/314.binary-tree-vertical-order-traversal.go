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
	val, row int
}
type s314 struct {
	order      map[int][]item
	minC, maxC int
}

func (s *s314) dfs(node *TreeNode, r, c int) {
	if node == nil {
		return
	}
	s.maxC = max(s.maxC, c)
	s.minC = min(s.minC, c)
	s.order[c] = append(s.order[c], item{
		val: node.Val,
		row: r,
	})
	s.dfs(node.Left, r+1, c-1)
	s.dfs(node.Right, r+1, c+1)
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	s := s314{order: map[int][]item{}}
	s.dfs(root, 0, 0)

	var res [][]int
	for c := s.minC; c <= s.maxC; c++ {
		curCol := s.order[c]

		sort.SliceStable(curCol, func(i, j int) bool {
			return curCol[i].row < curCol[j].row
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

