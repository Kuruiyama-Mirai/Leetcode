package leetcode

/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */

// @lc code=start
type Difference struct {
	diff []int
}

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	df := NewDifference(nums)

	for _, trip := range trips {
		val := trip[0]
		i := trip[1]
		j := trip[2] - 1
		df.Increment(i, j, val)
	}
	res := df.Result()

	for i := 0; i < len(res); i++ {
		if res[i] > capacity {
			return false
		}
	}
	return true
}
func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = d.diff[i] + res[i-1]
	}
	return res
}

// @lc code=end
