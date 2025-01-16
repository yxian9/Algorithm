package leetcode

import "slices"

// @leet start
type UnionFind struct {
	tree map[int]int // map[node]root
}

func (u *UnionFind) find(node int) int {
	if root, ok := u.tree[node]; ok {
		if root == node {
			return root
		}
		parent := u.find(root)
		u.tree[node] = parent
		return parent
	}
	return node
}

func (u *UnionFind) union(node1, node2 int) {
	// n1 and n2 are from same tree
	// compress and join connection
	r1, r2 := u.find(node1), u.find(node2)
	u.tree[r1] = r2
}

func (u *UnionFind) init() {
	u.tree = map[int]int{}
}

func accountsMerge(accounts [][]string) [][]string {
	var uf UnionFind
	uf.init()
	emailToTreeID := map[string]int{}
	// build tree graph, each list consider as a tree, with idx as tree id
	for idx, info := range accounts {
		for _, email := range info[1:] {
			// email is the only unique indentifier
			if treeID, ok := emailToTreeID[email]; ok {
				// these two tree should merge
				// because same email show in both tree
				uf.union(treeID, idx)
			} else {
				// new association setted for email, idx
				// not necessary saved in the uf, only union when overlap found
				emailToTreeID[email] = idx
			}
		}
	}
	emailGroups := map[int][]string{}
	for email, treeID := range emailToTreeID {
		rootID := uf.find(treeID)
		emailGroups[rootID] = append(emailGroups[rootID], email)
	}

	var ret [][]string
	for rootID, emails := range emailGroups {
		slices.Sort(emails)
		temp := []string{accounts[rootID][0]}
		temp = append(temp, emails...)
		ret = append(ret, temp)
	}

	return ret
}

// @leet end

