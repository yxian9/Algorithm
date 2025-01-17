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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	var ret int
	var dfs func(x *TreeNode) bool
	var pre *TreeNode
	dfs = func(x *TreeNode) bool {
		if x == nil {
			return false
		}
		l := dfs(x.Left)
		if l {
			return true
		}
		if float64(x.Val) >= target {
			if pre == nil {
				ret = x.Val
			} else {
				ldist := target - float64(pre.Val)
				rdist := float64(x.Val) - target
				// if ldist == rdist {
				// 	ret = pre.Val
				// } else if ldist <= rdist {
				if ldist <= rdist {
					ret = pre.Val
				} else {
					ret = x.Val
				}

			}
			return true
		}
		pre = x
		r := dfs(x.Right)
		return r
	}
	if dfs(root) {
		return ret
	}
	return pre.Val
}

// @leet end

func closestValue2(root *TreeNode, target float64) int {
	var ret *int
	var pre *TreeNode
	var dfs func(x *TreeNode) bool
	dfs = func(x *TreeNode) bool {
		if x == nil {
			return false
		}
		l := dfs(x.Left)
		if l {
			return true
		}
		if pre != nil {
			if float64(pre.Val) <= target && target <= float64(x.Val) {
				// will not work. not necessary place in between
				ldist := target - float64(pre.Val)
				rdist := float64(x.Val) - target
				if ldist < rdist {
					ret = &(pre.Val)
				} else {
					ret = &(x.Val)
				}
				return true
			}
		}
		pre = x
		if ret == nil {
			ret = &(x.Val)
		}
		r := dfs(x.Right)
		return r
	}
	dfs(root)
	return *ret
}

