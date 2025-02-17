# @leet start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        l = 0
        maxL = 0
        freq = [0] * 128

        def valid():
            for i in freq:
                if i > 1:
                    return False
            return True

        for r, c in enumerate(s):
            freq[ord(c)] += 1
            while not valid():
                freq[ord(s[l])] -= 1
                l += 1
            maxL = max(maxL, r - l + 1)

        return maxL


# @leet end

