package leetcode

// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		dummy = &ListNode{}
		cur   = dummy
	)
	for l1 != nil || l2 != nil || carry == 1 {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		total := n1 + n2 + carry
		carry = total / 10
		// assign before move. avoid nil // build
		cur.Next = &ListNode{Val: total % 10} // NOTE: add to next not current
		cur = cur.Next
	}
	return dummy.Next
}

// @leet end

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		head  = &ListNode{}
		cur   = head
	)
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		total := n1 + n2 + carry
		carry = total / 10
		if cur == nil {
			cur = &ListNode{}
		}
		cur.Val = total % 10
		if l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
		// cur = cur.Next NOTE: not work, will break if not attach first
	}
	if carry == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return head
}

