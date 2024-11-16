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
func dfsbuild2(postEnd, inStart, length int, postorder []int, inorderM map[int]int) *TreeNode {
	if length == 0 {
		return nil
	}
	rootVal := postorder[postEnd]
	inorderNeedle := inorderM[rootVal]
	leftSize := inorderNeedle - inStart
	rightSize := length - leftSize - 1

	l := dfsbuild2(postEnd-1-rightSize, inStart, leftSize, postorder, inorderM)
	r := dfsbuild2(postEnd-1, inorderNeedle+1, rightSize, postorder, inorderM)
	newNode := &TreeNode{
		Val:   rootVal,
		Left:  l,
		Right: r,
	}
	return newNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderM := make(map[int]int)
	for i, v := range inorder {
		inorderM[v] = i
	}
	newTree := dfsbuild2(len(postorder)-1, 0, len(inorder), postorder, inorderM)
	return newTree
}

// @leet end

