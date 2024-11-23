// @leet start
function findTargetSumWays(nums: number[], target: number): number {
  const memo = new Map<string, number>();
  function dfs(level: number, cur: number): number {
    const key = `${level}:${cur}`;
    if (memo.has(key)) return memo.get(key)!;
    if (level == nums.length) {
      return target == cur ? 1 : 0;
    }
    const l = dfs(level + 1, cur + nums[level]);
    const r = dfs(level + 1, cur - nums[level]);
    memo.set(key, l + r);
    return l + r;
  }
  return dfs(0, 0);
}
// @leet end
function findTargetSumWays2(nums: number[], target: number): number {
  let sum = 0;
  for (const i of nums) {
    sum += i;
  }
  if (Math.abs(target) > sum) return 0;
  if ((target + sum) % 2 == 1) return 0;
  const total = Math.floor((target + sum) / 2);
  const dp = Array(total + 1).fill(0);
  dp[0] = 1;
  for (const item of nums) {
    for (let j = dp.length - 1; j >= item; j--) {
      dp[j] += dp[j - item];
    }
  }
  return dp[dp.length - 1];
}

