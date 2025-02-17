from typing import Optional
from type import ListNode


# @leet start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        dummy = ListNode()
        cur = dummy
        carry = 0
        while l1 is not None or l2 is not None or carry == 1:
            n1 = n2 = 0
            if l1 is not None:
                n1 = l1.val
                l1 = l1.next
            if l2 is not None:
                n2 = l2.val
                l2 = l2.next
            total = n1 + n2 + carry
            cur.next = ListNode(total % 10)
            carry = total // 10
            cur = cur.next

        return dummy.next


# @leet end

