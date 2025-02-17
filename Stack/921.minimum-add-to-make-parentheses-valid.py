# @leet start
class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        st = []
        for c in s:
            if c == "(":
                st.append(c)
            else:
                if st and st[-1] == "(":
                    st.pop()
                else:
                    st.append(c)
        return len(st)


# @leet end

