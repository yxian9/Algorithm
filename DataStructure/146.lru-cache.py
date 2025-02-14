# @leet start


class n:
    def __init__(self, key=0, val=0, prev=None, next=None):
        self.key = key
        self.val = val
        self.prev = prev
        self.next = next


class LRUCache:
    def __init__(self, capacity: int):
        dummy = n()
        dummy.next = dummy
        dummy.prev = dummy
        self.capacity = capacity
        self.dummy = dummy
        self.m: dict[int, n] = {}

    def get(self, key: int) -> int:
        if key in self.m:
            x = self.m[key]
            self.detach(x)
            self.appendLeft(x)
            return x.val
        else:
            return -1

    def detach(self, x: n):
        x.prev.next = x.next  # type: ignore
        x.next.prev = x.prev  # type: ignore
        x.prev, x.next = None, None

    def appendLeft(self, x: n):
        x.prev = self.dummy
        x.next = self.dummy.next
        x.prev.next = x
        x.next.prev = x  # type: ignore

    def put(self, key: int, value: int) -> None:
        if key in self.m:
            x = self.m[key]
            x.val = value
            self.detach(x)
            self.appendLeft(x)
            return
        nn = n(key, value)
        self.m[key] = nn
        self.appendLeft(nn)
        if len(self.m) > self.capacity:
            lastItem = self.dummy.prev
            if lastItem:  # for type annotation
                self.detach(lastItem)
                del self.m[lastItem.key]


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
# @leet end

