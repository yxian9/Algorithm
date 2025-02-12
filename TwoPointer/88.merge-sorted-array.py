from typing import List


# @leet start
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        p1, p2 = m - 1, n - 1
        for i in range(m + n - 1, -1, -1):
            match True:
                case _ if p2 < 0:
                    break
                case _ if p1 < 0:
                    for j in range(p2 + 1):
                        nums1[j] = nums2[j]
                    break
                case _ if nums2[p2] >= nums1[p1]:
                    nums1[i] = nums2[p2]
                    p2 -= 1
                case _:
                    nums1[i] = nums1[p1]
                    p1 -= 1


# @leet end

