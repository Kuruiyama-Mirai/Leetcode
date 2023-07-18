package leetcode

//链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

//二叉树定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//116完美二叉树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//429多叉树
type ChildrenNode struct {
	Val      int
	Children []*Node
}

//取最大值
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//取最小值
func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
