package leetcode

// @leet start
 func distanceK(root *TreeNode, target *TreeNode, k int) []int {
   var (
     childToParen = map[*TreeNode]*TreeNode{}
     res          []int
     queue        = []*TreeNode{target}
     seen         = map[*TreeNode]bool{target: true}
     step         int
   )
   var dfs func(child, paren *TreeNode)
   dfs = func(child, paren *TreeNode) {
     if child == nil {
       return
     }
     childToParen[child] = paren
     dfs(child.Left, child)
     dfs(child.Right, child)
   }
   dfs(root, nil)
   // bfs to get target node
   for len(queue) > 0 {
     if step == k {
       for _, v := range queue {
         res = append(res, v.Val)
       }
       return res
     }
     for range len(queue) {
       lh := queue[0]
       queue = queue[1:]
       for _, item := range [3]*TreeNode{lh.Left, lh.Right, childToParen[lh]} {
         if item != nil && !seen[item] {
           queue = append(queue, item)
           seen[item] = true
         }
       }
     }
     step++
   }

   return res
 }

func distanceK2(root *TreeNode, target *TreeNode, K int) []int {
	childParMap := map[*TreeNode]*TreeNode{}

	// DFS annotates the parent
	var dfs func(child, par *TreeNode)
	dfs = func(child, par *TreeNode) {


	q := []*TreeNode{target}
	seen := map[int]bool{}
	seen[target.Val] = true
	var dist int
	var res []int = []int{}

	// BFS find the nodes at distance k from target node.
	for len(q) > 0 {
		if dist == K {
			for _, v := range q {
				res = append(res, v.Val)
			}
			return res
		}
		dist++
		length := len(q)
		for range length {
			head := q[0]
			q = q[1:]
			// if seen[head.Val] {
			// 	continue
			// }
			if head.Left != nil && !seen[head.Left.Val] {
				// all bsf seen happend during append step !!
				seen[head.Left.Val] = true
				q = append(q, head.Left)
			}
			if head.Right != nil && !seen[head.Right.Val] {
				seen[head.Right.Val] = true
				q = append(q, head.Right)
			}
			parNode := childParMap[head]
			if parNode != nil && !seen[parNode.Val] {
				seen[parNode.Val] = true
				q = append(q, parNode)
			}
		}
	}
	return []int{}
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK_simpleVersion(root *TreeNode, target *TreeNode, k int) []int {
	var (
		childToParen = map[*TreeNode]*TreeNode{}
		res          = []int{}
		queue        = []*TreeNode{target}
		seen         = map[*TreeNode]bool{}
		dist         int
	)
	// build childParen relation
	var dfs func(child, paren *TreeNode)
	dfs = func(child, paren *TreeNode) {
		if child == nil {
			return
		}
		childToParen[child] = paren
		dfs(child.Left, child)
		dfs(child.Right, child)
	}
	dfs(root, nil)
	// bfs
	for len(queue) > 0 {
		if dist == k {
			for _, v := range queue {
				if seen[v] || v == nil {
					continue
				}
				res = append(res, v.Val)
			}
			break
		}
		dist++
		for range len(queue) {
			lh := queue[0]
			queue = queue[1:]
			if seen[lh] || lh == nil { // way clean. need pay attention here
				continue // especially compare with 958. the coninue does not work
			}
			seen[lh] = true
			queue = append(queue, lh.Left)
			queue = append(queue, lh.Right)
			queue = append(queue, childToParen[lh])
		}

	}
	return res
}
