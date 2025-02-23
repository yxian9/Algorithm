// @leet start
/**
 * @param {number[][]} points
 * @param {number} k
 * @return {number[][]}
 */
class HeapItem {
  constructor(item, priority = item) {
    this.item = item;
    this.priority = priority;
  }
}

class MinHeap {
  constructor() {
    this.heap = [];
  }
  swap(a, b) {
    const temp = this.heap[a];
    this.heap[a] = this.heap[b];
    this.heap[b] = temp;
  }

  push(node) {
    this.heap.push(node);
    this.bubbleUp();
  }

  bubbleUp() {
    let idx = this.heap.length - 1;

    while (idx > 0) {
      const pIdx = Math.floor((idx - 1) / 2);

      if (this.heap[pIdx].priority < this.heap[idx].priority) break;
      // if the parent is bigger than the child then swap the parent and child
      this.swap(idx, pIdx);
      idx = pIdx;
    }
  }

  pop() {
    const min = this.heap[0];
    this.heap[0] = this.heap[this.size() - 1];
    this.heap.pop();
    this.bubbleDown();
    return min;
  }

  bubbleDown() {
    let idx = 0;
    let cur = 0;
    const n = this.heap.length;

    while (idx < n) {
      const l = 2 * idx + 1;
      const r = l + 1;

      if (l < n && this.heap[cur].priority > this.heap[l].priority) {
        cur = l;
      } // NOTE: can not use break here
      if (r < n && this.heap[cur].priority > this.heap[r].priority) {
        cur = r;
      }
      if (cur === idx) break;
      this.swap(cur, idx);
      idx = cur;
    }
  }

  peek() {
    return this.heap[0];
  }

  size() {
    return this.heap.length;
  }
}

function kClosest(points, k) {
  function dist(point) {
    return -(point[0] ** 2 + point[1] ** 2);
  }

  const maxHeap = new MinHeap();

  for (let i = 0; i < points.length; i++) {
    const pt = points[i];
    if (i < k) {
      maxHeap.push(new HeapItem(pt, dist(pt)));
    } else {
      maxHeap.push(new HeapItem(pt, dist(pt)));
      maxHeap.pop();
    }
  }

  const res = [];
  for (let i = 0; i < k; i++) {
    res.push(maxHeap.pop().item);
  }
  return res;
}

// @leet end

