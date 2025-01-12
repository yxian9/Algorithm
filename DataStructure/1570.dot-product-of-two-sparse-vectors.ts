// @leet start
class SparseVector {
  nonZero = new Map<number, number>();
  constructor(nums: number[]) {
    for (const [idx, v] of nums.entries()) {
      if (v !== 0) {
        this.nonZero.set(idx, v);
      }
    }
  }

  // Return the dotProduct of two sparse vectors
  dotProduct(vec: SparseVector): number {
    let res = 0;
    for (const [k, v] of this.nonZero) {
      res += v * (vec.nonZero.get(k) ?? 0);
    }
    return res;
  }
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * var v1 = new SparseVector(nums1)
 * var v2 = new SparseVector(nums1)
 * var ans = v1.dotProduct(v2)
 */
// @leet end

