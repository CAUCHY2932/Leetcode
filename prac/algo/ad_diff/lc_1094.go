package ad_diff

func carPooling(trips [][]int, capacity int) bool {
    // diff arr plus start end plus
    diff := make([]int, 1001)
    for i :=0; i < len(trips); i++ {
        p, s, e := trips[i][0], trips[i][1], trips[i][2]
        diff[s] += p
        diff[e] -= p
    }
    nums := make([]int, 10001)
    nums[0] = diff[0]
    for i := 1; i < 1001; i++ {
        nums[i] = nums[i-1] + diff[i]
        if nums[i] > capacity {
            return false
        }
    }
    return true
}
