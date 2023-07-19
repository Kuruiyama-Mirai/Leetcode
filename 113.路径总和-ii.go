package leetcode

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	traverse113(root, &res, new([]int), targetSum)
	return res
}

func traverse113(root *TreeNode, result *[][]int, curPath *[]int, targetSum int) {
	if root == nil {
		return
	}
	targetSum -= root.Val
	*curPath = append(*curPath, root.Val)

	//找到叶子节点了
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		pathCopy := make([]int, len(*curPath))
		copy(pathCopy, *curPath)
		*result = append(*result, pathCopy)
	}
	traverse113(root.Left, result, curPath, targetSum)
	traverse113(root.Right, result, curPath, targetSum)

	*curPath = (*curPath)[:len(*curPath)-1]
}

// @lc code=end
