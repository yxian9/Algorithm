from typing import List

def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        ret = [0] * n
        st = []
        for i in range(n - 1, -1, -1):
            # NOTE: dirction is opposite to (next) greater/smaller
            # when do compare, put the need update item to left side
            # pop invalid item, put need check item on left hand, <= is invalid
            while st and temperatures[st[-1]] <= temperatures[i]:
                st.pop()

            if st:
                ret[i] = st[-1] - i
            else:
                ret[i] = 0
            st.append(i)
        return ret
    
# @leet start
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        st = []

        for i in range(n - 1, -1, -1):
            v = temperatures[i]
            # NOTE: maintain  num[i]  -> <- 1 2 3
            if len(st) == 0:
                st.append(i)
                continue

            if v < temperatures[st[-1]]:  # when this condition, answer is ready
                res[i] = st[-1] - i
                st.append(i)  #
                continue

            while st and v >= temperatures[st[-1]]:  # until v < st[-1]
                st.pop()

            if st:
                res[i] = st[-1] - i
            st.append(i)  # put next large 9o

        return res


# @leet end
class Solution2:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        st = []

        for i in range(n - 1, -1, -1):
            v = temperatures[i]
            # NOTE: maintain  <- 1 2 3
            while st and v >= temperatures[st[-1]]:
                st.pop()
            if st:
                res[i] = st[-1] - i
            st.append(i)  # put next large 9o

        return res


class Solution3:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        st = []

        for i in range(n):
            v = temperatures[i]
            # NOTE: maintain  -> 1 2 3

            while st and v > temperatures[st[-1]]:
                # if v <= st[-1] v can not update any v is the smallest
                res[st[-1]] = i - st[-1]
                st.pop()

            st.append(i)  # temp in stac does not found large number
            # NOTE: first met if the first result
            # put not found next large into stack

        return res


class Solution4:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        st = []

        for i, t in enumerate(temperatures):
            # 3,2,1 -> <- num[i]
            if len(st) == 0 or t <= temperatures[st[-1]]:
                st.append(i)
                continue
            ## t > st[-1] and keep doing if ok
            while st and t > temperatures[st[-1]]:
                idx = st.pop()
                res[idx] = i - idx
            st.append(i)  # can not add first
        return res
