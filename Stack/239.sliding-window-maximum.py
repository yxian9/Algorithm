from collections import deque
from typing import List


# @leet start
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        res, stack = [], deque()

        def push(v):
            while len(stack) > 0 and v > stack[-1]:
                stack.pop()
            stack.append(v)

        def pop(v):
            if len(stack) > 0 and v == stack[0]:
                stack.popleft()

        def update():
            res.append(stack[0])

        for i, v in enumerate(nums):
            if i == k:
                break
            push(v)
        update()
        for i in range(k, len(nums)):
            push(nums[i])
            pop(nums[i - k])
            update()
        return res


# @leet end

