# @leet start
class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        if len(s) % 2 == 1:
            return False
        sema = 0
        for i in range(len(s)):
            if s[i] == "(" or locked[i] == "0":
                sema += 1
            else:
                sema -= 1
            if sema < 0:
                return False
        sema = 0
        for i in range(len(s) - 1, -1, -1):
            if s[i] == ")" or locked[i] == "0":
                sema += 1
            else:
                sema -= 1
            if sema < 0:
                return False
        return True


# @leet end

