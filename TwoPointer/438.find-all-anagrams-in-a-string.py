from typing import List


# @leet start
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        n1, n2 = len(s), len(p)
        if n2 > n1:
            return []
        freq = [0] * 128
        ret = []
        for c in p:
            freq[ord(c)] += 1

        def valid():
            return all(i == 0 for i in freq)

        for i in range(n2):
            c = s[i]
            freq[ord(c)] -= 1
        if valid():
            ret.append(0)

        for i in range(n2, n1):
            c = s[i]
            d = s[i - n2]
            freq[ord(c)] -= 1
            freq[ord(d)] += 1
            if valid():
                ret.append(i - n2 + 1)
        return ret


# @leet end

