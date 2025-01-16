package leetcode

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return merge2(lists[0], lists[1])
	default:
		mid := len(lists) / 2
		l := mergeKLists(lists[:mid])
		r := mergeKLists(lists[mid:])
		return merge2(l, r)
	}
}

func merge2(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		a.Next = merge2(a.Next, b)
		return a
	} else {
		b.Next = merge2(a, b.Next)
		return b
	}
}

// @leet end

