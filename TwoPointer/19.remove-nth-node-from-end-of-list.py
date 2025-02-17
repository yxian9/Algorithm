from typing import Optional, cast
from type import ListNode


# @leet start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        # fast = slow = dummy.next # NOTE: dummy head effetc
        fast = slow = dummy
        for _ in range(n):
            fast = fast.next

        while fast.Next is not None:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next

        return dummy.next


# @leet end

