package company

/*
有n个物品，每个物品有人个属性，第件物品的第个属性用个正整数表示记为4，两个不同的物品 ，
被称为是完美对的当且仅当ai1 十ai1 -a2+@52 -.·--a+aik，求完美对的个数。
*/

//用一个全局变量接收 不然每次回溯就会被覆盖掉
var res = 0

func isPerfect(nums1, nums2 []int) bool {
	temp := nums1[0] + nums2[0]
	for i := 1; i < len(nums1); i++ {
		if nums1[i]+nums2[i] != temp {
			return false
		}
	}
	return true
}

func backtrackingfutu(nums [][]int, path [][]int, index int) {
	if len(path) == 2 {
		if isPerfect(path[0], path[1]) {
			res++
		}
		return
	}
	for i := index; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrackingfutu(nums, path, index+1)
		path = path[:len(path)-1]
	}
}
