from typing import List


# @leet start
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        i = len(digits) - 1
        for i in range(i, -1, -1):
            if digits[i] + 1 != 10:
                digits[i] += 1
                return digits
            digits[i] = 0
            # digits[i - 1] += 1 # NOTE: double dipple
        if i == 0:
            return [1] + digits
        else:
            return digits


# @leet end

