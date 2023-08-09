package leetcode

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	//1.先从左到右，当右边的大于左边的就加1
	//2.再从右到左，当左边的大于右边的就再加1
	need := make([]int, len(ratings))
	sum := 0
	//优先保证每人至少1颗糖
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	//从左到右
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			//相邻的要大
			need[i+1] = need[i] + 1
		}
	}
	//从右到左
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			need[i] = max(need[i], need[i+1]+1)
		}
	}
	for i := range need {
		sum += need[i]
	}
	return sum
}

// @lc code=end
