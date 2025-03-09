# @leet start
class node:
    def __init__(self, key, value, next=None, prev=None) -> None:
        self.value = value
        self.key = key
        self.next = next
        self.prev = prev

    def detach(self):
        self.prev.next = self.next  # type: ignore
        self.next.prev = self.prev  # type: ignore
        self.prev = self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        dummy = node(0, 0)
        dummy.next = dummy.prev = dummy
        self.dummy = dummy
        self.capacity = capacity
        self.m = {}

    def appendLeft(self, x):
        x.prev = self.dummy
        x.next = self.dummy.next
        x.prev.next = x
        x.next.prev = x  # type: ignore

    def getNode(self, key):
        if key in self.m:
            x = self.m[key]
            x.detach()
            self.appendLeft(x)
            return x
        return None

    def get(self, key: int) -> int:
        n = self.getNode(key)
        if n:
            return n.value
        return -1

    def put(self, key: int, value: int) -> None:
        n = self.getNode(key)
        if n:
            n.value = value
            return

        if len(self.m) == self.capacity:
            tail = self.dummy.prev
            tail.detach()  # type: ignore
            del self.m[tail.key]  # type: ignore

        nn = node(key, value)
        self.m[key] = nn
        self.appendLeft(nn)

