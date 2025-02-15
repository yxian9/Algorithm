# @leet start
class Solution:
    def isValid(self, s: str) -> bool:
        map = {
            "(": ")",
            "{": "}",
            "[": "]",
        }
        st = []
        for c in s:
            if c in map:
                st.append(map[c])
            else:
                if len(st) == 0 or st.pop() != c:
                    return False
        return len(st) == 0


# @leet end

