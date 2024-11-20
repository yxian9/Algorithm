// @leet start
function combinationSum2(candidates: number[], target: number): number[][] {
  candidates.sort();
  const res: number[][] = [];
  const path: number[] = [];
  const n = candidates.length;
  function dfs(idx: number, total: number) {
    if (total >= target) {
      if (total == target) {
        res.push(path.slice());
      }
      return;
    }
    for (let i = 0; i < n; i++) {
      if (i > idx && candidates[i] === candidates[i - 1]) {
        continue;
      }
      const v = candidates[i];
      path.push(v);
      dfs(i + 1, total + v);
      path.pop();
    }
  }
  dfs(0, 0);
  return res;
}
// @leet end

