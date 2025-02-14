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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root.Left, root.Right}

	for len(stack) > 0 {
		l, r := stack[0], stack[1]
		stack = stack[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		// if l == nil || r == nil {
		// 	return l ==r // not work!!
		// }
		if l.Val != r.Val {
			return false
		}
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

// @leet end
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(Left, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}
	if Left == nil || Right == nil {
		return false
	}
	if Left.Val != Right.Val {
		return false
	}
	return isSame(Left.Left, Right.Right) && isSame(Left.Right, Right.Left)
	// mirror shape
	// not return isSame(Left.Left, Left.Left) && isSame(Right.Left, Right.Right)
}

func isSymmetric_wrong(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}
	// this not work. only consider single branch.
	// need to track both side!!
	return isSymmetric(root.Left) && isSymmetric(root.Right)
}

