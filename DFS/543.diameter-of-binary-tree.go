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
type sol543 struct {
	ans int
}

func (s *sol543) maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := s.maxDepth(node.Left)
	right := s.maxDepth(node.Right)
	s.ans = max(s.ans, left+right)
	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	solution := sol543{}
	solution.maxDepth(root)
	return solution.ans
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

// 🦊 its maxDepth from certain level!
func (s *sol543) maxDepth(node *TreeNode, step int) int {
	if node == nil {
		return step
	}
	left := s.maxDepth(node.Left, step+1)
	right := s.maxDepth(node.Right, step+1)
	return max(left, right)
}

func (s *sol543) travel(node *TreeNode) {
	if node == nil {
		return
	}
	//⏰ once put status into args, it can not be memory!!
	//🦊 need practice more on the backtravelse
	left := s.maxDepth(node.Left, 0)
	right := s.maxDepth(node.Right, 0)
	s.ans = max(left+right, s.ans)
	// need travel twice !!
	s.travel(node.Left)
	s.travel(node.Right)
}

func diameterOfBinaryTree(root *TreeNode) int {
	solution := sol543{}
	solution.travel(root)
	return solution.ans
}

