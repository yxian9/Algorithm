# @leet start


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        window = [0] * 128
        for c in s1:
            window[ord(c)] += 1

        def valid():
            return all(i == 0 for i in window)

        for r in range(len(s1)):
            c = s2[r]
            window[ord(c)] -= 1
        if valid():
            return True

        for r in range(len(s1), len(s2)):
            c = s2[r]
            d = s2[r - len(s1)]
            window[ord(c)] -= 1
            window[ord(d)] += 1

            if valid():
                return True

        return False


# @leet end

