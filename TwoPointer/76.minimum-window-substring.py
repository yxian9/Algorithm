# @leet start
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        freq = [0] * 128
        count = len(t)
        minLen = bestL = l = 0
        for c in t:
            freq[ord(c)] += 1
        ## go over s
        for r, c in enumerate(s):
            freq[ord(c)] -= 1  # NOTE: can be negative even for freq count
            # NOTE: freq like presSUm, always update.count is different
            if freq[ord(c)] >= 0:  # make sure each site is 0
                count -= 1

            while count == 0:
                curL = r - l + 1
                if minLen == 0 or curL < minLen:
                    minLen = curL
                    bestL = l
                l_c = s[l]
                l += 1
                freq[ord(l_c)] += 1  # NOTE: + not -
                if freq[ord(l_c)] > 0:
                    count += 1

        return s[bestL : bestL + minLen]


# @leet end
def minWindow(self, s: str, t: str) -> str:
        window = Counter()
        need = Counter(t)
        l = r = minL = bestL = valid = 0
        for r in range(len(s)):
            c = s[r]

            if need[c]:
                window[c] += 1
                if window[c] == need[c]:
                    valid += 1

            while valid == len(
                need
            ):  # NOTE: can include other char, no need fix window
                curL = r - l + 1
                if minL == 0 or curL < minL:
                    minL = curL
                    bestL = l
                d = s[l]
                l += 1

                if d in need:
                    window[d] -= 1
                    if window[d] < need[d]:
                        valid -= 1

        return s[bestL : bestL + minL]
