// @leet start
function combinationSum(candidates: number[], target: number): number[][] {
  const res: number[][] = [];
  const path: number[] = [];
  function dfs(total: number, i: number) {
    if (total >= target) {
      if (total === target) {
        res.push([...path]);
      }
      return;
    }
    if (i === candidates.length) return;
    dfs(total, i + 1);
    const v = candidates[i];
    path.push(v);
    dfs(total + v, i);
    path.pop();
  }

  dfs(0, 0);
  return res;
}
// @leet end
function combinationSum2(candidates: number[], target: number): number[][] {
  const res: number[][] = [];
  const path: number[] = [];
  function dfs(total: number, idx: number) {
    if (total >= target) {
      if (total === target) {
        res.push([...path]);
      }
      return;
    }
    for (let i = idx; i < candidates.length; i++) {
      const v = candidates[i];
      path.push(v);
      dfs(total + v, i);
      path.pop();
    }
  }

  dfs(0, 0);
  return res;
}

