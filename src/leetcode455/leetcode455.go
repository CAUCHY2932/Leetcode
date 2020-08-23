package leetcode455

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if g[gi] <= s[si] {
            gi++
        }
        si++
    }
    return gi
}
