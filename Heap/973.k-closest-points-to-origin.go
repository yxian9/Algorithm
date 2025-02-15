package leetcode

import "container/heap"

// @leet start
type maxHeap [][]int

func dist(x []int) int { return x[0]*x[0] + x[1]*x[1] }

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool { return dist(h[i]) > dist(h[j]) }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *maxHeap) Pop() (x any) {
	x = (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	var h maxHeap
	heap.Init(&h)
	for _, point := range points {
		if len(h) < k {
			heap.Push(&h, point)
			continue
		}
		if dist(point) < dist(h[0]) {
			heap.Push(&h, point)
			heap.Pop(&h)
		}
	}
	return h
}

// @leet end

func kClosest(points [][]int, k int) [][]int {
	var h maxHeap
	heap.Init(&h)
	for _, point := range points {
		if len(h) < k { // checked before push
			heap.Push(&h, point)
			continue
		}
		heap.Push(&h, point)
		heap.Pop(&h)

	}
	return h
}

// NOTE:
func kClosest(points [][]int, k int) [][]int {
	var h maxHeap
	heap.Init(&h)
	for _, point := range points {
		heap.Push(&h, point)
		if len(h) > k { // checked after push.  clean up
			heap.Pop(&h)
		} // no code after this line
	}
	return h
}
