from typing import List


# @leet start
class Solution:
    def jump(self, nums: List[int]) -> int:
        ret = curR = nextR = 0
        # greed not dp, we can chose any step not fixed step
        for i in range(len(nums) - 1):
            nextR = max(nextR, i + nums[i])
            if curR == i:
                curR = nextR
                ret += 1

        return ret


# @leet end

