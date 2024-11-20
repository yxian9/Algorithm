# @leet start
class Solution:
    def maxPal(self, s: str, l: int, r: int, cur: str) -> str:
        res = ""
        while l >= 0 and r < len(s) and s[l] == s[r]:
            res = s[l : r + 1]
            l -= 1
            r += 1
        return res if len(res) > len(cur) else cur

    def longestPalindrome(self, s: str) -> str:
        res = ""
        for i in range(len(s)):
            res = self.maxPal(s, i, i, res)
            res = self.maxPal(s, i, i + 1, res)
        return res


# @leet end
