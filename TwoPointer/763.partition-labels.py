from typing import List, Optional
from collections import defaultdict


# @leet start
class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        lastIdex = {}
        for i, c in enumerate(s):
            lastIdex[c] = i

        ret = []
        size = end = 0

        for i, c in enumerate(s):
            end = max(lastIdex[c], end)
            size += 1
            if i == end:
                ret.append(size)
                size = end = 0

        return ret


# @leet end

