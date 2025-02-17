from type import Node
from typing import Optional
from collections import deque

# @leet start
"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""
"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""


class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        m = {}  # NOTE: simple map

        def dfs(node):
            if node is None:
                return
            m[node] = Node(node.val)
            for item in node.neighbors:
                if item not in m:
                    dfs(item)
                m[node].neighbors.append(m[item])

        dfs(node)
        if node in m:
            return m[node]
        return None




class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        if node is None:
            return None
        old_new: dict[Node, Node] = {}
        old_new[node] = Node(node.val)
        q: deque[Node] = deque([node])
        while len(q) > 0:
            cur = q.popleft()
            cur_copy = old_new[cur]
            for item in cur.neighbors:
                if item not in old_new:
                    q.append(item)  # NOTE: to append. the order may be different
                    nn = Node(item.val)
                    old_new[item] = nn
                copy_item = old_new[item]
                cur_copy.neighbors.append(copy_item)

        return old_new[node]


# @leet end


class Solution1:  # NOTE: self connected. timeout
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        if node is None:
            return None
        nn = Node(node.val)
        for item in node.neighbors:
            nn.neighbors.append(self.cloneGraph(item))
        return nn


class Solution2:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        if node is None:
            return None
        copyList: list[None | Node] = [None] * 101
        # head = Node(node.val)
        # copyList[head.val] = head

        def dfs(n: Node):
            # if n is None:
            #     return
            nn = Node(n.val)
            copyList[n.val] = nn
            for item in n.neighbors:  # NOTE: item will not be nil!!
                # neighber can be empy list
                if copyList[item.val] is None:
                    dfs(item)
                nn.neighbors.append(copyList[item.val])  # type: ignore

        dfs(node)
        return copyList[node.val]
        # return head  # NOTE: note work first try will overwrite the head instance
