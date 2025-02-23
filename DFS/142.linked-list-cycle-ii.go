package leetcode

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	s, f := head, head
	for {
		if f.Next == nil {
			return nil
		}
		s, f = s.Next, f.Next.Next
		if s == f {
			f = head
			for {
				if s == f {
					return s
				}
				s, f = s.Next, f.Next
			}
		}

	}
}

// @leet end

