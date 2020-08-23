package leetcode455

import (
	"reflect"
	"testing"
)

func assertWithParameters(input [][]int, output int, t *testing.T) {
	got := findContentChildren(input[0], input[1])
	if !reflect.DeepEqual(got, output) {
		t.Errorf("input: %v want: %v but got:%v", input, output, got)
	}
}

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		input  [][]int
		output int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{1, 1},
			},
			output: 1,
		},
		{
			input: [][]int{
				{1, 2},
				{1, 2, 3},
			},
			output: 2,
		},
	}
	for _, c := range cases {
		assertWithParameters(c.input, c.output, t)
	}

}
