// @leet start
function countSubstrings(s: string): number {
  let ret = 0;
  function check(l: number, r: number) {
    while (l >= 0 && r < s.length && s[l] === s[r]) {
      if (s[l] == s[r]) {
        ret++;
      }
      l--;
      r++;
    }
  }
  for (let i = 0; i < s.length; i++) {
    check(i, i);
    check(i, i + 1);
  }

  return ret;
}
// @leet end

