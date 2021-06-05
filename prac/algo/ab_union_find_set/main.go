package ab_union_find_set

var parent []int

func union(from, to int) {
	parent[parent[from]] = find(to)
}

func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func initUnionFind(arr []int) {
	parent = make([]int, len(arr))
	for k := range arr {
		parent[k] = k
	}
}


