from typing import List, Optional


# @leet start
class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        ret = [-1, -1]
        if right <= 2:
            return ret

        def getPrim():
            isPrim = [True] * (right + 1)
            isPrim[0] = isPrim[1] = False
            for n in range(2, int(right**0.5) + 1):
                if isPrim[n]:
                    for j in range(n * n, right + 1, n):
                        isPrim[j] = False
            return [i for i, v in enumerate(isPrim) if v and i >= left]

        prims = getPrim()
        if len(prims) <= 1:
            return ret
        diff = None
        for i in range(1, len(prims)):
            curDiff = prims[i] - prims[i - 1]
            if diff is None or curDiff < diff:
                ret = [prims[i - 1], prims[i]]
                diff = curDiff
        return ret


# @leet end

