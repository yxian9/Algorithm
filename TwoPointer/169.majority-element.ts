// @leet start
function majorityElement(nums: number[]): number {
  let major = nums[0];
  let count = 0;
  for (const v of nums) {
    if (v === major) {
      count++;
    } else {
      count--;
    }
    if (count < 0) {
      count = 0;
      major = v;
    }
  }
  return major;
}
// @leet end

