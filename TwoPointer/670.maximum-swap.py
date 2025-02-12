# @leet start
class Solution:
    def maximumSwap(self, num: int) -> int:
        bestLoc = [0] * 10
        for i, v in enumerate(str(num)):
            bestLoc[int(v)] = i
        #
        numList = list(str(num))
        #
        for i, curInt in enumerate(numList):
            for v in range(9, int(curInt), -1):
                if bestLoc[v] > i:
                    numList[i], numList[bestLoc[v]] = numList[bestLoc[v]], numList[i]
                    return int("".join(numList))
        return num


# @leet end

