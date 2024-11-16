from typing import List


# @leet start
class Solution:
    def __init__(self) -> None:
        self.res = []
        self.mapping = "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"

    def dfs(self, digits: str, path: list[str]):
        if len(path) == len(digits):
            self.res.append("".join(path))
            return
        for c in self.mapping[int(digits[len(path)])]:
            path.append(c)
            self.dfs(digits, path)
            path.pop()

    def letterCombinations(self, digits: str) -> List[str]:
        if digits == "":
            return []
        self.dfs(digits, [])
        return self.res


# @leet end
