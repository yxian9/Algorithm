from collections import deque
from typing import List

def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        ret = []
        q = deque()

        def push(v):
            while q and q[-1] < v:
                q.pop()
            q.append(v)

        def pop(v):
            if q[0] == v:
                q.popleft()

        def update():
            ret.append(q[0])

        for i in range(k):
            push(nums[i])
        update()
        for i in range(k, len(nums)):
            push(nums[i])
            pop(nums[i - k])
            update()

        # for i in range(len(nums)):
        #     if i < k - 1:
        #         push(nums[i])
        #     else:
        #         push(nums[i])
        #         update()
        #         pop(nums[i - k + 1])
        return ret
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

