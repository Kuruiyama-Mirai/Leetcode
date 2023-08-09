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
