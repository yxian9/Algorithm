# @leet start
class Solution:
    def calculate(self, s: str) -> int:
        ret = prev = num = 0
        op = "+"
        for c in s + "+":
            if c == " ":
                continue
            if c.isdigit():
                num = num * 10 + int(c)
                continue
            match op:
                case "+":
                    prev = num
                case "-":
                    prev = -num
                case "*":
                    ret -= prev
                    prev *= num
                case "/":
                    ret -= prev
                    prev = int(prev / num)
            ret += prev
            op, num = c, 0
        return ret


# @leet end
