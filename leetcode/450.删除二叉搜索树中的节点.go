package leetcode

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 这两个 if 把情况 1 和 2 都正确处理了
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 处理情况 3
		// 获得右子树最小的节点
		minNode := getMin(root.Right)
		// 删除右子树最小的节点
		root.Right = deleteNode(root.Right, minNode.Val)
		// 用右子树最小的节点替换 root 节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(node *TreeNode) *TreeNode {
	// BST 最左边的就是最小的
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// @lc code=end
