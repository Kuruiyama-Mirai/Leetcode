package leetcode

/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 */

// @lc code=start
type DifferenceV1 struct {
	diff []int
}

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	df := NewDifferenceV1(nums)

	for _, book := range bookings {
		i := book[0] - 1
		j := book[1] - 1
		val := book[2]
		df.Increment(i, j, val)
	}
	return df.Result()
}

func NewDifferenceV1(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

func (d *DifferenceV1) IncrementV1(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *DifferenceV1) ResultV1() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = d.diff[i] + res[i-1]
	}
	return res
}

// @lc code=end
