import { ListNode } from "./type";
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

function walk2(pre: ListNode | null, curr: ListNode | null): ListNode | null {
  if (!curr) return pre;
  let temp: ListNode | null = curr.next; // hold
  curr.next = pre; // break
  pre = curr; // move
  curr = temp; // move 2
  return walk2(pre, curr); // do not forget return
}

function reverseList(head: ListNode | null): ListNode | null {
  return walk2(null, head);
}
// @leet end

