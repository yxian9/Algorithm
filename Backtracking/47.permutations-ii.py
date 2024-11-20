from typing import List


# @leet start
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        n = len(nums)
        seen = [False] * n
        res, path = [], []

        def dfs():
            if len(path) == n:
                res.append(path[:])
                return
            for i, v in enumerate(nums):
                if seen[i]:
                    continue
                ## this is different compared with the idx version
                # must use seen[] bool
                # in the combination, idx is included, so not from 0, exclude all duplicated
                if i > 0 and nums[i] == nums[i - 1] and seen[i]:
                    continue
                seen[i] = True
                path.append(v)
                dfs()
                path.pop()
                seen[i] = False

        dfs()
        return res


# @leet end
