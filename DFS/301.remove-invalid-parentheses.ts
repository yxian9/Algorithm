// @leet start
/**
 * @param {string} s
 * @return {string[]}
 */
var removeInvalidParentheses = function (s) {
  const index = [];
  let arr = [s];
  const se = new Set();
  for (let i = 0; i < s.length; i++) {
    if (s[i] == ")" && index.length && s[index[index.length - 1]] == "(") {
      index.pop();
    } else if (s[i] == "(" || s[i] == ")") {
      index.push(i);
    }
  }

  while (index.length) {
    if (s[index[index.length - 1]] == ")") break;
    const a = [], b = index.pop();
    for (const j of arr) {
      for (let i = b; i < j.length; i++) {
        if (j[i] == "(") {
          const string_ = j.slice(0, i) + j.slice(i + 1);
          if (!se.has(string_)) {
            se.add(string_);
            a.push(string_);
          }
        }
      }
    }
    arr = a;
  }
  for (let k = 0; k < index.length; k++) {
    let a = [];
    for (let j of arr) {
      for (let i = 0; i <= index[k] - k; i++) {
        if (j[i] == ")") {
          let string_ = j.slice(0, i) + j.slice(i + 1);
          if (!se.has(string_)) {
            se.add(string_);
            a.push(string_);
          }
        }
      }
    }
    arr = a;
  }
  return arr;
};
// @leet end

