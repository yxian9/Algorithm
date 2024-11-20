// @leet start
function subsetsWithDup(nums: number[]): number[][] {
  const res: number[][] = [];
  const path: number[] = [];
  function dfs(start: number) {
    res.push(path.slice());
    for (let i = start; i < nums.length; i++) {
      if (i > start && nums[i] === nums[i - 1]) {
        continue;
      }
      const v = nums[i];
      path.push(v);
      dfs(i + 1);
      path.pop();
    }
  }
  dfs(0);
  return res;
}
// @leet end
