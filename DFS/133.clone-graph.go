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
func cloneGraph(node *Node) *Node {
	var (
		copied = map[*Node]*Node{}
		dfs    func(*Node)
	)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		copied[node] = &Node{Val: node.Val}
		for _, item := range node.Neighbors {
			if copied[item] == nil {
				dfs(item)
			}
			copied[node].Neighbors = append(copied[node].Neighbors, copied[item])
		}
	}
	dfs(node)
	return copied[node]
}

// @leet end

func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	nn := &Node{Val: node.Val}
	for _, v := range node.Neighbors {
		nn.Neighbors = append(nn.Neighbors, cloneGraph(v))
		// NOTE: this will not work seem it will self connected.
	}
	return nn
}

