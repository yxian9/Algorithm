from typing import List


# @leet start
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        n, path, res = len(nums), [], []
        nums.sort()

        def dfs(start):
            res.append(path[:])
            for i in range(start, n):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                path.append(nums[i])
                dfs(i + 1)
                path.pop()

        dfs(0)
        return res


# @leet end

