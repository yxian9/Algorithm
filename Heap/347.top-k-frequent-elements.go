package leetcode

import "container/heap"

// @leet start
type f struct {
	v, count int
}
type minHeap []f

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(f)) }
func (h *minHeap) Pop() (x any) {
	n := h.Len()
	x = (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	var h minHeap
	heap.Init(&h)
	for v, freq := range count {
		if len(h) < k {
			heap.Push(&h, f{v, freq})
			continue
		}
		if freq > h[0].count {
			heap.Pop(&h)
			heap.Push(&h, f{v, freq})
		}
	}

	var res []int
	for _, v := range h {
		res = append(res, v.v)
	}
	return res
}

// @leet end

