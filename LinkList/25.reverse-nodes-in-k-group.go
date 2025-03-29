package leetcode

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		dummy  = &ListNode{Next: head}
		preGrp = dummy
		getNth = func(cur *ListNode, k int) *ListNode {
			for cur != nil && k > 0 {
				cur = cur.Next
				k--
			}
			return cur
		}
	)
	for {
		kth := getNth(preGrp, k)
		if kth == nil {
			break
		}
		nextGrp := kth.Next
		prev, cur := nextGrp, preGrp.Next
		for cur != nextGrp {
			temp := cur.Next
			cur.Next = prev
			prev = cur
			cur = temp
		}
		preGrp, preGrp.Next = preGrp.Next, kth
		// temp := preGrp.Next
		// preGrp.Next = kth
		// preGrp = temp
	}
	return dummy.Next
}

// @leet end

