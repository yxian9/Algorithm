from typing import Optional, Self


class TreeNode:
    def __init__(
        self, val=0, left: Optional[Self] = None, right: Optional[Self] = None
    ):
        self.val = val
        self.left = left
        self.right = right


class Node:
    def __init__(
        self,
        val: int,
        next: Optional[Self] = None,
        random: Optional["Node"] = None,
        neighbors: list[Self] = [],
    ):
        self.val = int(val)
        self.next = next
        self.random = random
        self.neighbors = neighbors if neighbors is not None else []


class ListNode:
    def __init__(self, val=0, next: Self | None = None):
        self.val = val
        self.next = next
