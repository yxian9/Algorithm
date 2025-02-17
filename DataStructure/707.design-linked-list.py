# @leet start
from dataclasses import dataclass
from typing import Optional, Self, cast


@dataclass(slots=True)
class nd:
    val: int = 0
    prev: Self | None = None
    next: Self | None = None


class MyLinkedList:
    def __init__(self):
        dummy = nd()
        dummy.next = dummy.prev = dummy
        self.dummy = dummy
        self.length = 0

    def get_n(self, index: int) -> Optional[nd]:
        if index >= self.length:
            return None
        cur = cast(nd, self.dummy.next)
        for _ in range(index):
            cur = cast(nd, cur.next)
        return cur

    def get(self, index: int) -> int:
        n = self.get_n(index)
        if n is None:
            return -1
        return n.val

    def insert(self, x: nd, val: int) -> None:
        nn = nd(val)
        nn.prev = x.prev
        nn.next = x
        cast(nd, nn.prev).next = nn
        nn.next.prev = nn
        self.length += 1

    def addAtHead(self, val: int) -> None:
        self.insert(cast(nd, self.dummy.next), val)

    def addAtTail(self, val: int) -> None:
        self.insert(self.dummy, val)

    def addAtIndex(self, index: int, val: int) -> None:
        if index == self.length:
            self.addAtTail(val)
            return
        x = self.get_n(index)
        if x is not None:
            self.insert(x, val)

    def deleteAtIndex(self, index: int) -> None:
        x = self.get_n(index)
        if x is not None:
            cast(nd, x.prev).next = x.next
            cast(nd, x.next).prev = x.prev
            x.prev = x.next = None
            self.length -= 1


# Your MyLinkedList object will be instantiated and called as such:
# obj = MyLinkedList()
# param_1 = obj.get(index)
# obj.addAtHead(val)
# obj.addAtTail(val)
# obj.addAtIndex(index,val)
# obj.deleteAtIndex(index)
# @leet end

