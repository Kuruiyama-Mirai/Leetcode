package leetcode

/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	left, right := 0, len(arr)-1

	for left < len(arr)-1 && arr[left] < arr[left+1] {
		left++
	}
	for right > 0 && arr[right] < arr[right-1] {
		right--
	}
	//走到转折点要判断一下
	if left == right && left != 0 && right != len(arr)-1 {
		return true
	}
	return false
}

// @lc code=end
