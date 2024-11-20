// @leet start
function twoSum(nums: number[], target: number): number[] {
  const m = new Map<number, number>()
  for( const [idx, v] of nums.entries()){
    const diff = target - v
    if( m.has(diff)){
      return [idx, m.get(diff)!]
    }
    m.set(v, idx)
  }
  return []
};
// @leet end
