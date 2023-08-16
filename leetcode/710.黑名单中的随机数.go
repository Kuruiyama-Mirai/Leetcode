package leetcode

import "math/rand"

/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 */

// @lc code=start
type Solution710 struct {
	sz     int //白名单数字
	valMap map[int]int
}

// 构造函数，初始化黑名单映射表
// N: [0, N) 中的数字
// blacklist: 黑名单中的数字
func Constructor710(n int, blacklist []int) Solution710 {
	s := Solution710{
		sz:     n - len(blacklist),
		valMap: make(map[int]int),
	}
	for _, b := range blacklist {
		s.valMap[b] = 666
	}
	last := n - 1
	for _, b := range blacklist {
		// 已经在区间 [sz, N) 的数字可以直接忽略
		if b >= s.sz {
			continue
		}
		// 寻找一个不在黑名单中的位置 last
		for _, ok := s.valMap[last]; ok; _, ok = s.valMap[last] {
			last--
		}
		s.valMap[b] = last
		last--
	}
	return s
}

func (s *Solution710) Pick() int {
	// 随机选取一个索引
	index := rand.Intn(s.sz)
	// 如果 index 命中了黑名单中的数，
	// 则返回它被映射到的那个位置
	if val, ok := s.valMap[index]; ok {
		return val
	}
	// 如果 index 没命中黑名单，则直接返回
	return index

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end
