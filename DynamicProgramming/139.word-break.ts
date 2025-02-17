// @leet start
function wordBreak(s: string, wordDict: string[]): boolean {
  const choice = new Set<number>();
  const m = new Set<string>();
  for (const v of wordDict) {
    m.add(v);
    choice.add(v.length);
  }
  const dp = new Array(s.length + 1).fill(false);
  dp[0] = true;
  for (let i = 0; i <= s.length; i++) {
    for (const step of choice) {
      if (i >= step && dp[i - step]) {
        if (m.has(s.slice(i - step, i))) {
          dp[i] = true;
          break;
        }
      }
    }
  }

  return dp[s.length];
}

// @leet end

