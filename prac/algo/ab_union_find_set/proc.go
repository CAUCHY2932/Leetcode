package ab_union_find_set

func findCircleNum(isConnected [][]int) (ans int) {


	n := len(isConnected)
	parent = make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	// var find func(int)int
	find := func(i int) int {
		for parent[i] != i {
			i = parent[i]
		}
		return i
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		for j := i +1;j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}

	ret := 0
	
	for i, p := range parent {
		if i == p {
			ret++
		}
	}
	return ret



    // n := len(isConnected)
    // parent := make([]int, n)
    // for i := range parent {
    //     parent[i] = i
    // }
    // var find func(int) int
    // find = func(x int) int {
    //     if parent[x] != x {
    //         parent[x] = find(parent[x])
    //     }
    //     return parent[x]
    // }
    // union := func(from, to int) {
    //     parent[find(from)] = find(to)
    // }

    // for i, row := range isConnected {
    //     for j := i + 1; j < n; j++ {
    //         if row[j] == 1 {
    //             union(i, j)
    //         }
    //     }
    // }
    // for i, p := range parent {
    //     if i == p {
    //         ans++
    //     }
    // }
    // return
}
