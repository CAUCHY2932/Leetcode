package leetcode198

import (
	"reflect"
	"testing"
)

func assertWithParameters(input []int, output int, t *testing.T) {
	got := rob(input)
	if !reflect.DeepEqual(got, output) {
		t.Errorf("input: %v want: %v but got:%v", input, output, got)
	}
}

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 3, 1},
			output: 4,
		},
		{
			input:  []int{2, 7, 9, 3, 1},
			output: 12,
		},
	}
	for _, c := range cases {
		assertWithParameters(c.input, c.output, t)
	}

}
