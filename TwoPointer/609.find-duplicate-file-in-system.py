from typing import List
from collections import defaultdict


# @leet start
class Solution:
    def findDuplicate(self, paths: List[str]) -> List[List[str]]:
        ret = []
        m = defaultdict(list)
        for path in paths:
            dir, *files = path.split()
            for file in files:
                fn, content = file.split("(")
                m[content].append(f"{dir}/{fn}")

        for v in m.values():
            if len(v) > 1:
                ret.append(v)
        return ret


# @leet end

