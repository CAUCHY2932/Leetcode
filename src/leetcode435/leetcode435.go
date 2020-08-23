package leetcode435

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals)==0{
		return 0
	}
	res := 0
	var interSort slcArr = intervals
	sort.Sort(interSort)
	fmt.Println(interSort)
	cur := interSort[0][1]
	for _, v := range interSort[1:]{
		if v[0] < cur{
			cur = min(cur, v[1])
			res++
		}else{
			cur = v[1]
		}
	}
	return res
}


type slcArr [][]int


func (s slcArr) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}
func (s slcArr) Len() int {
	return len(s)

}

func (s slcArr)Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func min(x, y int)int{
	if x <y {
		return x
	}
	return y
}

// type slc struct {
// 	pre   int
// 	after int
// }

// type slcArr []slc


// func (s slcArr) Less(i, j int) bool {
// 	return s[i].after < s[j].after
// }
// func (s slcArr) Len() int {
// 	return len(s)

// }

// func (s slcArr)Swap(i, j int) {
// 	s[i].after, s[j].after = s[j].after, s[i].after
// }
