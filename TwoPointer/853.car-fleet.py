from typing import List, Optional
from collections import defaultdict


# @leet start
class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        st = []
        pair = [(target - position[i], speed[i]) for i in range(len(speed))]
        pair.sort()
        for d, s in pair:
            st.append(d / s)
            if len(st) > 1 and st[-1] <= st[-2]:
                st.pop()
        return len(st)


# @leet end

