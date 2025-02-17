# @leet start
class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        l1, l2 = len(num1) - 1, len(num2) - 1
        carry = 0
        res = []
        while l1 >= 0 or l2 >= 0:
            n1 = n2 = 0
            if l1 >= 0:
                n1 = ord(num1[l1]) - ord("0")
            if l2 >= 0:
                n2 = ord(num2[l2]) - ord("0")

            total = n1 + n2 + carry
            res.append(str(total % 10))
            carry = total // 10
            l1 -= 1
            l2 -= 1
        if carry == 1:
            res.append("1")
        res.reverse()
        return "".join(res)


# @leet end

