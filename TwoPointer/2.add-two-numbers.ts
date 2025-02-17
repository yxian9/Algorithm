import { ListNode } from "./type.ts";

// @leet start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function addTwoNumbers(
  l1: ListNode | null,
  l2: ListNode | null,
): ListNode | null {
  const dummy = new ListNode();
  let cur = dummy;
  let carry = 0;
  while (l1 != null || l2 != null || carry === 1) {
    let n1 = 0;
    let n2 = 0;
    if (l1 != null) {
      n1 = l1.val;
      l1 = l1.next;
    }
    if (l2 != null) {
      n2 = l2.val;
      l2 = l2.next;
    }
    const total = n1 + n2 + carry;
    carry = Math.floor(total / 10);
    cur.next = new ListNode(total % 10);
    cur = cur.next;
  }

  return dummy.next;
}
// @leet end

