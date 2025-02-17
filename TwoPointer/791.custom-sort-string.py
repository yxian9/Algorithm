# @leet start
class Solution:
    def customSortString(self, order: str, s: str) -> str:
        freq = [0] * 128
        res = ""
        for c in s:
            freq[ord(c)] += 1
        for c in order:
            res += c * freq[ord(c)]
            freq[ord(c)] = 0
        for i, count in enumerate(freq):
            res += chr(i) * count
        return res


# @leet end

