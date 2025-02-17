from collections import deque
from typing import List


# @leet start
class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        all_letter = [chr(i + ord("a")) for i in range(26)]

        def possible(s: str):
            for i in range(len(s)):
                for c in all_letter:
                    # NOTE: s[len(s):] will return "" without key error
                    yield s[:i] + c + s[i + 1 :]

        map = set(wordList)
        # if endWord not in map:
        #     return 0
        step = 1
        q = deque([beginWord])
        while len(q) > 0:
            for _ in range(len(q)):
                top = q.popleft()
                if top == endWord:
                    return step
                for var in possible(top):
                    if var in map:
                        q.append(var)
                        map.remove(var)
            step += 1
        return 0


# @leet end

