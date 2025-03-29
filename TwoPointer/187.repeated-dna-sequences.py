from typing import List, Optional
from collections import defaultdict, deque


# @leet start
class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        if len(s) < 10:
            return []
        q = deque()
        ans: set[str] = set()
        seen = set()

        def update():
            cur = "".join(q)
            if cur in seen:
                ans.add(cur)
            else:
                seen.add(cur)

        for i in range(10):
            q.append(s[i])
        update()
        for i in range(10, len(s)):
            q.popleft()
            q.append(s[i])
            update()

        return list(ans)


# @leet end

