from typing import List, Optional
from collections import defaultdict
from type import ListNode


# @leet start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(0, next=head)
        preGroup = dummy

        while True:
            kth = self.getNth(preGroup, k)
            if not kth:
                break

            nextGroup = kth.next

            prev, cur = nextGroup, preGroup.next
            while (
                cur != nextGroup
            ):  # NOTE: not cur != kth.next. kth will be update later
                temp = cur.next
                cur.next = prev
                prev = cur
                cur = temp

            temp = preGroup.next
            preGroup.next = kth
            preGroup = temp

        return dummy.next

    def getNth(self, cur, k):
        while cur and k > 0:
            cur = cur.next
            k -= 1
        return cur


# @leet end
