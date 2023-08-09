package leetcode

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob337(root *TreeNode) int {
	var robTree func(cur *TreeNode) []int
	robTree = func(cur *TreeNode) []int {
		//返回数组
		//下标为0记录不偷该节点所得到的的最大金钱
		//下标为1记录偷该节点所得到的的最大金钱
		if cur == nil {
			return []int{0, 0}
		}
		left := robTree(cur.Left)
		right := robTree(cur.Right)

		//偷当前节点
		robCur := cur.Val + left[0] + right[0]
		//不偷当前，最大值就是左右子树的值
		notRob := max(left[0], left[1]) + max(right[0], right[1])

		return []int{notRob, robCur}
	}
	res := robTree(root)
	return max(res[0], res[1])
}

// @lc code=end
