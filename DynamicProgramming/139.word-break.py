from typing import List


# @leet start
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        dp = [True] + [False] * len(s)
        choise = set([len(i) for i in wordDict])
        m = set(wordDict)
        for i in range(len(s) + 1):
            for step in choise:
                if i >= step and dp[i - step]:
                    if s[i - step : i] in m:
                        dp[i] = True
                        break

        return dp[len(s)]


# @leet end


class Solution2:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        memo = {}

        def dfs(i: int) -> bool:
            if i == len(s):
                return True
            if i > len(s):
                return False
            if i in memo:
                return memo[i]
            ret = False
            for item in wordDict:
                if s[i : i + len(item)] == item:
                    if dfs(i + len(item)):
                        ret = True
                        break
            memo[i] = ret
            return ret

        return dfs(0)
