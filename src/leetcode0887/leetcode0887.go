package leetcode0887

func superEggDrop(k int, n int) int {
	memo = make([][]int, k)
	for i := 0; i < k; i++ {
		memo[i] = make([]int, n)
	}
	return dp(k, n)
}

var memo [][]int

// k egg, n floor
func dp(k, n int) int {
	if memo[k][n] != 0 {
		return memo[k][n]
	}
	ans := 0
	if n == 0 {
		ans = 0
	} else if k == 1 {
		ans = n
	} else {
		lo := 1
		hi := n
		for lo+1 < hi {
			x := (lo + hi) / 2
			t1 := dp(k-1, x-1)
			t2 := dp(k, n-x)

			if t1 < t2 {
				lo = x
			} else if t1 > t2 {
				hi = x
			} else {
				hi = x
				lo = x

			}
			ans = 1 + min(max(dp(k-1, lo-1), dp(k, n-lo)), max(dp(k-1, hi-1), dp(k, n-hi)))
		}
		memo[k][n] = ans

	}
	return memo[k][n]
}

func max(x, y int) int {
	if x > y {

		return x
	}
	return y

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}
