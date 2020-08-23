package leetcode435

import (
	"reflect"
	"testing"
)

func assertWithParameters(input [][]int, output int, t *testing.T) {
	got := eraseOverlapIntervals(input)
	if !reflect.DeepEqual(output, got) {
		t.Errorf("input: %v want: %v but got: %v", input, output, got)
	}
}
func TestEraseOverlapIntervals(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output int
	}{
		{
			input: [][]int{
				{0, 1}, {3, 4}, {1, 2},
			},
			output: 0,
		},
		{},
	}
	for _, c := range testCases {
		assertWithParameters(c.input, c.output, t)
	}
}
