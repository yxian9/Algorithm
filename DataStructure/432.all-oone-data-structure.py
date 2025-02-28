# @leet start
class Node:
    def __init__(self, s, cnt, prev=None, next=None):
        self.s = {s}
        self.cnt = cnt
        self.prev = prev
        self.next = next

    def pop(self, key):
        self.s.remove(key)
        if len(self.s) == 0:
            self.prev.next = self.next  # type: ignore
            self.next.prev = self.prev  # type: ignore
            self.prev = self.next = None


class AllOne:
    def __init__(self):
        self.dummy = Node("", 0)
        self.dummy.prev = self.dummy
        self.dummy.next = self.dummy
        self.m = dict()

    def insert_before(self, nn, x):
        nn.prev = x.prev
        nn.next = x
        nn.prev.next = nn
        nn.next.prev = nn

    def insert_after(self, nn, x):
        nn.prev = x
        nn.next = x.next
        nn.prev.next = nn
        nn.next.prev = nn

    def inc(self, key: str) -> None:
        nn = None
        if key in self.m:
            cur = self.m[key]
            cnt = cur.cnt + 1
            if cnt == cur.next.cnt:
                nn = cur.next
                nn.s.add(key)
            else:
                nn = Node(key, cnt)
                self.insert_after(nn, cur)
            cur.pop(key)
        else:
            if self.dummy.next.cnt == 1:  # type: ignore  the first item may have 2 cnt
                nn = self.dummy.next
                nn.s.add(key)  # type: ignore
            else:
                nn = Node(key, 1)
                self.insert_before(nn, self.dummy.next)
        self.m[key] = nn

    def dec(self, key: str) -> None:
        if key not in self.m:
            return
        cur = self.m[key]
        if cur.cnt == 1:
            del self.m[key]
        else:
            cnt = cur.cnt - 1
            if cur.prev.cnt == cnt:
                nn = cur.prev
                nn.s.add(key)
                self.m[key] = nn
            else:
                nn = Node(key, cnt)
                self.insert_before(nn, cur)
                self.m[key] = nn
        cur.pop(key)

    def getMaxKey(self) -> str:  # type: ignore
        for s in self.dummy.prev.s:  # type: ignore
            return s

    def getMinKey(self) -> str:  # type: ignore
        for s in self.dummy.next.s:  # type: ignore
            return s


# Your AllOne object will be instantiated and called as such:
# obj = AllOne()
# obj.inc(key)
# obj.dec(key)
# param_3 = obj.getMaxKey()
# param_4 = obj.getMinKey()
# @leet end

