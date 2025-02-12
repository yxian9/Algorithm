# @leet start
class Solution:
    def simplifyPath(self, path: str) -> str:
        items = []
        for i in path.split("/"):
            match i:
                case "..":
                    if len(items) > 0:
                        items.pop()
                case "." | "":
                    continue
                case _:
                    items.append(i)

        return "/" + "/".join(items)


# @leet end
