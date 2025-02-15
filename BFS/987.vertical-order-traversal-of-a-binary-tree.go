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

func verticalTraversal(root *TreeNode) [][]int {
	var (
		l, r int
		ret  [][]int
	)
	mp := map[int][]node1{}
	queue := []node1{{root, 0, 0}}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		mp[top.col] = append(mp[top.col], top)
		l = min(l, top.col)
		r = max(r, top.col)
		if top.Left != nil {
			queue = append(queue, node1{top.Left, top.col - 1, top.row + 1})
		}
		if top.Right != nil {
			queue = append(queue, node1{top.Right, top.col + 1, top.row + 1})
		}
	}
	for i := l; i <= r; i++ {
		cur := mp[i]
		sort.Slice(cur, func(i, j int) bool {
			l, r := cur[i], cur[j]
			if l.row != r.row {
				return l.row < r.row
			}
			return l.TreeNode.Val < r.TreeNode.Val
		})
		temp := make([]int, len(cur))
		for i, v := range cur {
			temp[i] = v.TreeNode.Val
		}
		ret = append(ret, temp)
	}

	return ret
}

// @leet end

func verticalTraversal_2(root *TreeNode) [][]int {
	type item struct {
		c int
		n *TreeNode
	}
	type n struct {
		r, v int
	}
	var (
		minC, maxC, level int
		queue             = []item{{0, root}}
		m                 = map[int][]n{}
	)
	for len(queue) > 0 {

		for range len(queue) {
			cur := queue[0]
			queue = queue[1:]
			maxC, minC = max(cur.c, maxC), min(cur.c, minC)
			m[cur.c] = append(m[cur.c], n{level, cur.n.Val})
			if cur.n.Left != nil {
				queue = append(queue, item{cur.c - 1, cur.n.Left})
			}
			if cur.n.Right != nil {
				queue = append(queue, item{cur.c + 1, cur.n.Right})
			}
		}
		level++
	}
	ret := make([][]int, 0, maxC-minC+1)
	for i := minC; i <= maxC; i++ {
		// ret[i] = m[i] // have negative index
		curC := m[i]
		sort.Slice(curC, func(i, j int) bool {
			if curC[i].r == curC[j].r {
				return curC[i].v < curC[j].v
			}
			return curC[i].r < curC[j].r
		})
		temp := make([]int, len(curC))
		for i, v := range curC {
			temp[i] = v.v
		}
		ret = append(ret, temp)
	}
	return ret
}
