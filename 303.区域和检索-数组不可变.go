package leetcode /*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	nums []int
}

func ConstructorV1(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{nums: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
