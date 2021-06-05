package ac_slide_win

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
    diff := make([]int, len(s))
    for i := range s {
        diff[i] = int(abs(s[i], t[i]))
		fmt.Printf("get diff[i] %v, s[i] %v, t[i] %v\n", diff[i], s[i], t[i])
    }
    start := 0
    cost := 0
    maxLen := 0
    for end, d := range diff {
        cost += d
        for cost > maxCost {
            cost -= diff[start]
            start++
        }
        maxLen = max(maxLen, end - start + 1)
    }
    return maxLen
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x, y byte)byte {
    if x > y {
        return x - y
    }
    return y - x
}
