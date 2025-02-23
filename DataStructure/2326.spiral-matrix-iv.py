from typing import List, Self, cast
from typing import Optional

# @leet start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
from dataclasses import dataclass


@dataclass(frozen=True)
class pt:
    r: int
    c: int

    def mv(self, d: Self):
        return pt(self.r + d.r, self.c + d.c)


class Solution:
    def spiralMatrix(self, m: int, n: int, head: Optional[ListNode]) -> List[List[int]]:
        res = [[-1] * n for _ in range(m)]
        dirs = (
            pt(0, 1),
            pt(1, 0),
            pt(0, -1),
            pt(-1, 0),
        )

        def inside(p: pt):
            return 0 <= p.r < m and 0 <= p.c < n

        def gen():
            nonlocal head
            if head is None:
                return -1
            res = cast(int, head.val)
            head = head.next
            return res

        step = 0
        cur = pt(0, 0)
        while True:
            v = gen()
            if v == -1:
                return res
            res[cur.r][cur.c] = v
            np = cur.mv(dirs[step % 4])
            if not inside(np) or res[np.r][np.c] != -1:
                step += 1
                np = cur.mv(dirs[step % 4])
            cur = np
        # return res  # type: ignore


# @leet end
