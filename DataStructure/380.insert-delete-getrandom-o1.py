# @leet start
from random import randint


class RandomizedSet:
    def __init__(self):
        self.m = {}
        self.arr = []

    def insert(self, val: int) -> bool:
        if val in self.m:
            return False
        self.m[val] = len(self.arr)
        self.arr.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.m:
            return False

        idx = self.m[val]
        self.arr[idx] = self.arr[-1]
        self.m[self.arr[-1]] = idx

        self.arr.pop()
        self.m.pop(val)

        return True

    def getRandom(self) -> int:
        idx = randint(0, len(self.arr) - 1)
        return self.arr[idx]


# Your RandomizedSet object will be instantiated and called as such:
# obj = RandomizedSet()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
# @leet end
