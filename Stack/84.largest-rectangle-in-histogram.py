from typing import List


# @leet start
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = [-1]
        max_area = 0

        for i, h in enumerate(heights):
            while stack[-1] != -1 and h <= heights[stack[-1]]:
                height = heights[stack.pop()]
                width = i - stack[-1] - 1
                max_area = max(max_area, height * width)
            stack.append(i)

        while stack[-1] != -1:
            height = heights[stack.pop()]
            width = len(heights) - stack[-1] - 1
            max_area = max(max_area, height * width)

        return max_area

    def largestRectangleArea2(self, heights: List[int]) -> int:
        n = len(heights)
        l_small = [-1] * n
        r_small = [n] * n
        ret = 0
        st = []
        for i in range(n):
            while st and heights[st[-1]] >= heights[i]:
                st.pop()
            if st:
                l_small[i] = st[-1]
            else:
                l_small[i] = -1
            st.append(i)

        st.clear()
        for i in range(n - 1, -1, -1):
            while st and heights[st[-1]] >= heights[i]:
                st.pop()
            if st:
                r_small[i] = st[-1]
            else:
                r_small[i] = n
            st.append(i)

        for i in range(n):
            r, l = r_small[i], l_small[i]
            ret = max(ret, heights[i] * (r - l - 1))
        return ret


# @leet end
