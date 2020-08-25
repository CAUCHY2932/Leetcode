package leetcode70

import (
	"reflect"
	"testing"
)

func assertWithParameters(input int, output int, t *testing.T) {
	got := climbStairs(input)
	if !reflect.DeepEqual(got, output) {
		t.Errorf("input: %v want: %v but got:%v", input, output, got)
	}
}

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{
			input:  2,
			output: 2,
		},
		{
			input:  3,
			output: 3,
		},
	}
	for _, c := range cases {
		assertWithParameters(c.input, c.output, t)
	}

}
