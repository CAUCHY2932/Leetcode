package leetcode0304

type NumMatrix struct {
	rows   [][]int
	origin [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rowNum := len(matrix)
	coloumNum := len(matrix[0])

	rowPrefixs := make([][]int, rowNum)

	for i := 0; i < rowNum; i++ {

		rowPrefixs[i] = make([]int, coloumNum)

	}

	for i := 0; i < rowNum; i++ {
		for j := 0; j < coloumNum; j++ {
			if j == 0 {
				rowPrefixs[i][j] = rowPrefixs[i][0]
				continue
			}
			rowPrefixs[i][j] = rowPrefixs[i][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{
		rows:   rowPrefixs,
		origin: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for j := col1; j < col2; j++ {
		sum += this.rows[row2][j] - this.rows[row1][j] + this.origin[row1][j]

	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
