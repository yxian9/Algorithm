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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		lastItem := queue[length-1]
		res = append(res, lastItem.Val)
		for length > 0 {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			length--
		}
	}
	return res
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
func rightSideView(root *TreeNode) []int {
	var (
		res   []int
		queue = []*TreeNode{root}
	)
	// bsf
	for len(queue) > 0 {
		rh := queue[len(queue)-1]
		if rh != nil {
			res = append(res, rh.Val) // not work, some element is nil at end, need to make sure no nil get into queue
		}
		for range len(queue) {
			lh := queue[0]
			queue = queue[1:]
			if lh == nil {
				continue
			}
			queue = append(queue, lh.Left, lh.Right)
		}
	}

	return res
}

func rightSideView_clearstWay!!(root *TreeNode) []int {
	var (
		res   []int
		queue = []*TreeNode{root}
	)
	// bsf
	for len(queue) > 0 {
		for _, v := range queue {
			if v != nil {
				res = append(res, v.Val)
				break
			}
		}

		for range len(queue) {
			lh := queue[0]
			queue = queue[1:]
			if lh == nil {
				continue
			}
			queue = append(queue, lh.Right, lh.Left)
		}
	}

	return res
}


