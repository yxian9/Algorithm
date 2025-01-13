package leetcode

import "container/heap"

// @leet start
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() (x any) {
	x = (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := intHeap(nums[:k])
	heap.Init(&h)
	for _, v := range nums[k:] {
		if v > h[0] {
			heap.Pop(&h)
			heap.Push(&h, v)
		}
	}
	return h[0]
}

// @leet end
