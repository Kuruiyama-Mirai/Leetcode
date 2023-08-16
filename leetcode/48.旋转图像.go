package leetcode

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate48(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	for _, row := range matrix {
		reverseArr(row)
	}
}

func reverseArr(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
}

// @lc code=end
