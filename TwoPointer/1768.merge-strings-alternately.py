# @leet start
class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        n1, n2 = len(word1), len(word2)
        i1 = i2 = 0
        ret = []
        while i1 < n1 or i2 < n2:
            if i1 < n1 and i2 < n2:
                ret.append(word1[i1])
                ret.append(word2[i2])
                i1 += 1
                i2 += 1
                continue

            if i2 == n2:
                for i in range(i1, n1):
                    ret.append(word1[i])
                break
            if i1 == n1:
                for i in range(i2, n2):
                    ret.append(word2[i])
                break

        return "".join(ret)


# @leet end

