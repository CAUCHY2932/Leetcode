package ac_slide_win

import "testing"

type Input struct {
	s       string
	t       string
	maxCost int
}

type OutPut struct {
	maxLen int
}

func TestEqualSubstring(t *testing.T) {
	testCases := []struct {
		Give Input
		Want OutPut
	}{
		{
			Input{s: "abcd", t: "bcdf", maxCost: 3}, OutPut{maxLen: 3},
		},
	}
	for i, tc := range testCases {
		got := equalSubstring(tc.Give.s, tc.Give.t, tc.Give.maxCost)
		if got != tc.Want.maxLen {
			t.Errorf("[%d] want %v, but got %v", i, tc.Want.maxLen, got)
		}
	}

}
