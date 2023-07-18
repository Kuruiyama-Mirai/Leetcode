package leetcode

/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	left, right := 0, 1
	for _, w := range weights {
		left = Max(left, w)
		right += w
	}

	for left < right {
		mid := left + (right-left)/2
		if shipFunc(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 定义：当运载能力为 x 时，需要 f(x) 天运完所有货物,返回值即为需要的天数
func shipFunc(weights []int, x int) int {
	days := 0
	for i := 0; i < len(weights); {
		cap := x
		for i < len(weights) {
			if cap < weights[i] {
				break
			} else {
				cap -= weights[i]
				i++
			}
		}
		days++
	}
	return days
}

// @lc code=end
