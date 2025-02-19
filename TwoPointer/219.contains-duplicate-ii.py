from typing import List


# @leet start
class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        m = set()
        for i, v in enumerate(nums):
            if v in m:
                return True
            m.add(v)
            if i >= k:
                m.remove(nums[i - k])
        return False


# @leet end

