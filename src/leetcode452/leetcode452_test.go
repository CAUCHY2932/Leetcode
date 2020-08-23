package leetcode452

import (
	"reflect"
	"testing"
)

func assertWithParameters(input [][]int, output int, t *testing.T) {
	got := findMinArrowShots(input)
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
				{10,16},
				{2,8},
				{1,6},
				{7,12},
			},
			output: 2,
		},
	}
	for _, c := range cases {
		assertWithParameters(c.input, c.output, t)
	}

}
