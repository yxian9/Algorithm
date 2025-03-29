package leetcode

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	var (
		fast, slow                        = head, head
		first, second, temp, prev, t1, t2 *ListNode
	)
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second = slow.Next
	slow.Next = nil // NOTE: break into two

	for second != nil {
		temp = second.Next
		second.Next = prev
		prev = second
		second = temp
	}
	first, second = head, prev
	for second != nil {
		t1, t2 = first.Next, second.Next
		first.Next = second
		second.Next = t1
		first, second = t1, t2
	}
}

// @leet end

