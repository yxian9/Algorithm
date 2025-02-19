from typing import List


# @leet start
class Log:
    def __init__(self, s) -> None:
        id, status, time = s.split(":")
        self.t = int(time)
        self.id = int(id)
        self.typ = status


class Solution:
    def exclusiveTime(self, n: int, logs: List[str]) -> List[int]:
        ret = [0] * n  # tricky, interactive overlap ret and st

        st: list[Log] = []
        for s in logs:
            log = Log(s)
            if log.typ == "start":
                st.append(log)  # not update ret
            else:
                top = st.pop()
                lap = log.t - top.t + 1
                ret[top.id] += lap
                # NOTE: for different function, also need exclude overlap
                if st:
                    ret[st[-1].id] -= lap

        return ret


# @leet end

