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
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	x := root.Val
	if x > high {
		return rangeSumBST(root.Left, low, high)
	}
	if x < low {
		return rangeSumBST(root.Right, low, high)
	}

	return x + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// @leet end

func rangeSumBST2(root *TreeNode, low int, high int) int {
	var (
		st    = []*TreeNode{root}
		res   int
		valid = func(x int) bool { return x >= low && x <= high }
	)
	for len(st) > 0 {
		hd := st[0]
		// NOTE: st = st[:len(st)-1]  //incorret
		st = st[1:]
		if hd == nil {
			continue
		}
		if valid(hd.Val) {
			res += hd.Val
		}
		if hd.Val > low { // NOTE: not necessary in the first level, may be later level
			st = append(st, hd.Left)
		}
		if hd.Val < high {
			st = append(st, hd.Right)
		}
	}

	return res
}

