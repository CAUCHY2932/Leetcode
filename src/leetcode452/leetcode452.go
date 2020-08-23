package leetcode452

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	var pointsSort ballonsArrary = points
	sort.Sort(pointsSort)
	cnt := 1
	cur := pointsSort[0][1]
	for _, v := range pointsSort[1:] {
		if v[0] > cur {
			cnt++
			cur = v[1]
		}
	}
	return cnt
}

type ballonsArrary [][]int

func (b ballonsArrary) Len() int {
	return len(b)
}

func (b ballonsArrary) Less(i, j int) bool {
	return b[i][1] < b[j][1]
}

func (b ballonsArrary) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
