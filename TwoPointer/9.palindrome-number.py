# @leet start
class Solution:
    def isPalindrome(self, x: int) -> bool:
        num = x
        res = 0
        while num > 0:
            res = res * 10 + num % 10
            num //= 10
        return res == x


# @leet end

