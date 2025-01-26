package leetcode

// @leet start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	copies := [101]*Node{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil { // not necessary. since the range will exclude the nil node
			// eventually
			return
		}
		newNode := &Node{Val: node.Val}
		copies[node.Val] = newNode

		for _, neighbor := range node.Neighbors {
			if copies[neighbor.Val] == nil {
				dfs(neighbor)
			}
			newNode.Neighbors = append(newNode.Neighbors, copies[neighbor.Val])
		}
	}
	dfs(node)
	return copies[node.Val]
}

// @leet end

