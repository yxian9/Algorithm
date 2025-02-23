from typing import List


# @leet start
class Solution:
    def generate(self, n: int) -> List[List[int]]:
        res = [[1]]
        for i in range(1, n):
            x = [1]
            last = res[i - 1]
            for j in range(1, len(last)):
                x.append(last[j - 1] + last[j])
            x.append(1)
            res.append(x)

        return res


# @leet end

