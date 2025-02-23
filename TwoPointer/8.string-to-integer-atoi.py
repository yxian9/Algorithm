# @leet start
class Solution:
    def myAtoi(self, s: str) -> int:
        i, n = 0, len(s)
        sign = 1
        while i < n and s[i] == " ":
            i += 1
        if i < n and s[i] == "+" or s[i] == "-":
            if s[i] == "-":
                sign = -1
            i += 1
        num = 0
        while i < n and s[i].isdigit():
            num = num * 10 + ord(s[i]) - ord("0")
            i += 1
        num *= sign
        minR = -(2**31)
        maxR = 2**31 - 1

        if num > maxR:
            return maxR
        elif num < minR:
            return minR
        else:
            return num


# @leet end

