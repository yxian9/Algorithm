package leetcode

// @leet start
type top struct {
	res      int
	graph    map[int][]int
	indegree map[int]int
}

func (t *top) buildGraph(reqs [][]int) {
	for _, req := range reqs {
		t.graph[req[1]] = append(t.graph[req[1]], req[0])
	}
}

func (t *top) updateIndegree() {
	for _, chilren := range t.graph {
		for _, child := range chilren {
			t.indegree[child]++
		}
	}
}

func (t *top) topSort() {
	queue := []int{}
	for child, indegree := range t.indegree {
		if indegree == 0 {
			queue = append(queue, child)
		}
	}
	// start sort
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		t.res++
		// refill queue
		for _, child := range t.graph[parent] {
			t.indegree[child]--
			if t.indegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := map[int]int{}
	for i := range numCourses {
		indegree[i] = 0
	}
	s := top{
		res:      0,
		indegree: indegree,
		graph:    map[int][]int{},
	}
	s.buildGraph(prerequisites)
	s.updateIndegree()
	s.topSort()
	return s.res == numCourses
}

// @leet end

