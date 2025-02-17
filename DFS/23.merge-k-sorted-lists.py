from typing import List, Optional
from type import ListNode


# @leet start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        match len(lists):
            case 0:
                return None
            case 1:
                return lists[0]
            case _:
                mid = len(lists) // 2
                return self.merge2(
                    self.mergeKLists(lists[mid:]),
                    self.mergeKLists(lists[:mid]),
                )

    def merge2(self, a, b) -> Optional[ListNode]:
        if a is None:
            return b
        if b is None:
            return a
        if a.val < b.val:
            a.next = self.merge2(a.next, b)
            return a
        else:
            b.next = self.merge2(b.next, a)
            return b


# @leet end

