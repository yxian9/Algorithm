package leetcode

// @leet start
/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
	a, b := p, q
	for a != b {
		if a == nil {
			a = q
		} else {
			a = a.Parent
		}
		if b == nil {
			b = p
		} else {
			b = b.Parent
		}
	}
	return a
}

// @leet end

