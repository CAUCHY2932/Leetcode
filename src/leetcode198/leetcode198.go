package leetcode198

func rob(nums []int) int {

	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	for i := 2; i < length; i++ {
		dp[i] = maxValue(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[length-1]
}

func maxValue(x, y int) int {
	if x < y {
		return y
	}
	return x
}
