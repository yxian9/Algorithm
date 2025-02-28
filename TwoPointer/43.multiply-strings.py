# @leet start
class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        m = len(num1) - 1
        n = len(num2) - 1
        res = 0
        for i in range(m, -1, -1):
            x = num1[i]
            x = ord(x) - ord("0")
            temp = 0
            # for j, y in enumerate(nums2):
            for j in range(n, -1, -1):
                y = num2[j]
                y = ord(y) - ord("0")
                temp += pow(10, m - i) * x * pow(10, n - j) * y
            res += temp
        return str(res)


# @leet end

