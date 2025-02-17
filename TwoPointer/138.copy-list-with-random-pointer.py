from typing import Optional
from type import Node

# @leet start
"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""


class Solution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        cur = head
        m: dict[unknown:unknown] = {None: None}
        while cur is not None:
            m[cur] = Node(cur.val)
            cur = cur.next
        cur = head
        while cur is not None:
            m[cur].next = m[cur.next]
            m[cur].random = m[cur.random]
            cur = cur.next
        return m[head]


# @leet end

