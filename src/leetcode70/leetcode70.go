package leetcode70

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre, suffix := 1, 2
	for i := 2; i < n; i++ {
		pre, suffix = suffix, pre+suffix
	}
	return suffix
}
