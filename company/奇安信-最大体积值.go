/*
一个长方体，长宽高都是质数，已知长宽高之和为n【n为[6，10000]范围内的自然数。】，求这个长方体的体积最大值。
输入值：长宽高之和。
输出值：体积的最大可能值。
*/
package company

import "math"

func getMaxVolume(n int64) int64 {
	var res int64
	res = math.MinInt64
	var i int64
	nums := make([]int64, 0)
	for ; i <= n; i++ {
		if isPrime(i) {
			nums = append(nums, i)
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == n {
					temp := nums[i] * nums[j] * nums[k]
					res = max(res, temp)
				}
			}
		}
	}
	return res
}

func isPrime(n int64) bool {
	var i int64
	for ; i <= int64(math.Sqrt(float64(i))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
