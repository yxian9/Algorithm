# @leet start
from collections import defaultdict


class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        l = res = maxfreq = 0
        freq = defaultdict(int)

        for r, c in enumerate(s):
            freq[c] += 1
            maxfreq = max(maxfreq, freq[c])

            while r - l + 1 - maxfreq > k:
                freq[s[l]] -= 1
                l += 1

            res = max(res, r - l + 1)

        return res


# @leet end

