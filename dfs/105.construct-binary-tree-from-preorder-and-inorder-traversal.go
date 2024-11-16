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

func dfsbuild(preStart, inStart, length int, preorder []int, inorderM map[int]int) *TreeNode {
	if length == 0 {
		return nil
	}
	rootVal := preorder[preStart]
	inorderNeedle := inorderM[rootVal]
	leftSize := inorderNeedle - inStart
	rightSize := length - leftSize - 1

	l := dfsbuild(preStart+1, inStart, leftSize, preorder, inorderM)
	r := dfsbuild(preStart+1+leftSize, inorderNeedle+1, rightSize, preorder, inorderM)
	newNode := &TreeNode{
		Val:   rootVal,
		Left:  l,
		Right: r,
	}
	return newNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderM := make(map[int]int)
	for i, v := range inorder {
		inorderM[v] = i
	}
	newTree := dfsbuild(0, 0, len(inorder), preorder, inorderM)
	return newTree
}

// @leet end
