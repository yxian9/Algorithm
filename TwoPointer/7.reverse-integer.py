# @leet start
class Solution:
    def reverse(self, x: int) -> int:
        rev = x < 0
        if rev:
            x = -x
        ret = 0
        while x > 0:
            ret = ret * 10 + x % 10
            x //= 10
        if ret > 2**31 - 1:
            ret = 0
        return -ret if rev else ret


# @leet end

