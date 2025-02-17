from typing import List


# @leet start
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        strs.sort()
        lastItem, first = strs[-1], strs[0]
        i = 0
        for i in range(len(first)):
            if lastItem[i] != first[i]:
                # break # NOTE: is possible all match
                return first[:i]
        return first


# @leet end
