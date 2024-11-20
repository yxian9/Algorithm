// @leet start
function permuteUnique(nums: number[]): number[][] {
  nums.sort();
  const seen = nums.map(() => false);
  const res: number[][] = [];
  const path: number[] = [];
  function dfs(idx: number) {
    if (idx === nums.length) {
      res.push(path.slice());
      return;
    }
    for (const [i, v] of nums.entries()) {
      if (seen[i]) continue;
      if (i > 0 && nums[i] === nums[i - 1] && seen[i - 1]) continue;
      seen[i] = true;
      path[idx] = v;
      dfs(idx + 1);
      seen[i] = false;
    }
  }
  dfs(0);
  return res;
}
// @leet end

