# @leet start
from collections import Counter


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_c = Counter(s)
        for c in t:
            s_c.subtract(c)
        for i in s_c.values():
            if i != 0:
                return False
        return True


# @leet end

