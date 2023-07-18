package leetcode

import "math/rand"

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		nums:       []int{},
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	// 若 val 不存在，插入到 nums 尾部，
	// 并记录 val 对应的索引值
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	//若元素不存在，则无法删除，返回失败
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	index := this.valToIndex[val]
	this.valToIndex[this.nums[len(this.nums)-1]] = index
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]

	this.nums = this.nums[:len(this.nums)-1]
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
