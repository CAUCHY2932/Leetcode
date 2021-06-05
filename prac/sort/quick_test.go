package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 6, 5, 4, 2, 1, 3, 7, 8}
	quickSort(&arr)
	fmt.Println("result is ", arr)
	
}