# @leet start
class Solution:
    def decodeString(self, s: str) -> str:
        st = []
        for c in s:
            if c == "]":
                ct = patten = ""
                while st and st[-1] != "[":
                    patten = st.pop() + patten
                st.pop()
                while st and st[-1].isdigit():
                    ct = st.pop() + ct
                st.extend(int(ct) * patten)
            else:
                st.append(c)
        return "".join(st)


# @leet end

