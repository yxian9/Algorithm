from typing import List


# @leet start
class Solution:
    def validUtf8(self, data: List[int]) -> bool:
        remain = 0
        for byte in data:
            c = 0
            mask = 128  # 0b1000000

            while byte & mask != 0:
                c += 1
                mask = mask >> 1
                if c > 4:
                    return False

            if c != 1:
                if remain != 0:
                    return False
                if c > 1:
                    remain = c - 1
                continue

            if remain == 0:
                return False

            remain -= 1

        return remain == 0


# @leet end

