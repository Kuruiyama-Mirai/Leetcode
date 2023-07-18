package leetcode

import (
	"math/rand"
	"sort"
)

/*
 * @lc app=leetcode.cn id=528 lang=golang
 *
 * [528] 按权重随机选择
 */

// @lc code=start
type Solution528 struct {
	PreSum []int
}

// 构造函数 返回构造概率的前缀和数组
func Constructor528(w []int) Solution528 {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution528{
		PreSum: w,
	}
}

func (s *Solution528) PickIndex() int {
	x := rand.Intn(s.PreSum[len(s.PreSum)-1]) + 1
	return sort.SearchInts(s.PreSum, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end
