from typing import List
from collections import deque


# @leet start
class Solution:
    def valid(self, s):
        l = r = 0
        for c in s:
            match c:
                case "(":
                    l += 1
                case ")":
                    if l > 0:
                        l -= 1
                    else:
                        r += 1
        return l, r

    def removeInvalidParentheses(self, s: str) -> List[str]:
        res = []
        l, r = self.valid(s)

        def dfs(s, idx, l, r):
            if l == r == 0:
                if self.valid(s) == (0, 0):
                    res.append(s)
                    return
            for i in range(idx, len(s)):
                if i > 0 and s[i] == s[i - 1]:
                    continue
                if l > 0 and s[i] == "(":
                    n_s = s[:i] + s[i + 1 :]
                    dfs(n_s, i, l - 1, r)
                if r > 0 and s[i] == ")":
                    n_s = s[:i] + s[i + 1 :]
                    dfs(n_s, i, l, r - 1)

        dfs(s, 0, l, r)
        return res


# @leet end
class Solution2:
    def removeInvalidParentheses(self, s: str) -> List[str]:
        index = []
        arr = [s]
        seen = set()

        # Build a list of indices for the parentheses that are candidates for removal.
        for i, ch in enumerate(s):
            if ch == ")" and index and s[index[-1]] == "(":
                index.pop()
            elif ch in ("(", ")"):
                index.append(i)

        # Remove extra opening parentheses (from right side)
        while index:  # (
            if s[index[-1]] == ")":
                break
            new_arr = []
            b = index.pop()
            for string in arr:
                # Remove an opening parenthesis from position b onward
                for i in range(b, len(string)):
                    if string[i] == "(":
                        new_string = string[:i] + string[i + 1 :]
                        if new_string not in seen:
                            seen.add(new_string)
                            new_arr.append(new_string)
            arr = new_arr

        # Remove extra closing parentheses using the remaining indices.
        for k in range(len(index)):
            new_arr = []
            for string in arr:
                # Adjust index by subtracting the number of removals already done (k)
                for i in range(0, index[k] - k + 1):
                    if string[i] == ")":
                        new_string = string[:i] + string[i + 1 :]
                        if new_string not in seen:
                            seen.add(new_string)
                            new_arr.append(new_string)
            arr = new_arr

        return arr


class Solution:
    def check(self, s: str):
        l = r = 0
        for ch in s:
            if ch == "(":
                l += 1
            elif ch == ")":
                if l > 0:
                    l -= 1
                else:
                    r += 1
        return l, r

    def dfs(self, s: str, start: int, l: int, r: int):
        if l == 0 and r == 0:
            if self.check(s) == (0, 0):
                self.ans.append(s)
            return

        for i in range(start, len(s)):
            # Skip duplicate removals to avoid generating duplicate strings.
            if i > 0 and s[i] == s[i - 1]:
                continue

            if l > 0 and s[i] == "(":
                new_s = s[:i] + s[i + 1 :]
                self.dfs(new_s, i, l - 1, r)

            if r > 0 and s[i] == ")":
                new_s = s[:i] + s[i + 1 :]
                self.dfs(new_s, i, l, r - 1)

    def removeInvalidParentheses(self, s: str) -> List[str]:
        self.ans = []
        left, right = self.check(s)
        # Start the DFS from index 0 with the computed removal counts.
        self.dfs(s, 0, left, right)
        return self.ans

