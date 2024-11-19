// @leet start
function subsets(nums: number[]): number[][] {
  const res: number[][] = [];
  const path: number[] = [];
  function dfs(i: number) {
    res.push([...path]);
    for (let j = i; j < nums.length; j++) {
      path.push(nums[j]);
      dfs(j + 1);
      path.pop();
    }
  }
  dfs(0);
  return res;
}
// @leet end
function subsets2(nums: number[]): number[][] {
  const res: number[][] = [];
  const path: number[] = [];
  function dfs(i: number) {
    if (i === nums.length) {
      res.push([...path]);
      return;
    }
    dfs(i + 1);
    path.push(nums[i]);
    dfs(i + 1);
    path.pop();
  }
  dfs(0);
  return res;
}

