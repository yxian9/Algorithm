package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val    int
	Next   *Node
	Left   *Node
	Right  *Node
	Parent *Node
}
type ListNode struct {
	Val  int
	Next *ListNode
}
