// @leet start
function permute(nums: number[]): number[][] {
  const res: number[][] = [];
  const path: number[] = [];
  const seen: boolean[] = nums.map(() => false);
  function dfs(i: number) {
    if (i === nums.length) {
      res.push(path.slice());
      return;
    }
    for (const [idx, v] of nums.entries()) {
      if (seen[idx]) {
        continue;
      }
      seen[idx] = true;
      path[i] = v;
      dfs(i + 1);
      seen[idx] = false;
    }
  }
  dfs(0);
  return res;
}
// @leet end

