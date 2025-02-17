# @leet start
class Solution:
    def countSubstrings(self, s: str) -> int:
        ret = 0

        def search(l, r):
            while l >= 0 and r < len(s):
                if s[l] != s[r]:  # NOTE: need break
                    break
                if s[l] == s[r]:
                    nonlocal ret
                    ret += 1
                l -= 1
                r += 1

        for i in range(len(s)):
            search(i, i)
            search(i, i + 1)
        return ret


# @leet end

