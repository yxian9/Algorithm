# @leet start
class Solution:
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        if len(s1) != len(s2):
            return False
        temp = None
        firstTry = True
        for i in range(len(s1)):
            if s1[i] != s2[i]:
                if firstTry and temp is None:
                    temp = (s2[i], s1[i])
                elif firstTry and temp == (s1[i], s2[i]):
                    firstTry = False
                    temp = None
                else:
                    return False
        return temp is None


# @leet end

