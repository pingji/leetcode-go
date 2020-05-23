package problem0053

import (
	"testing"
)

func TestProblem0053(t *testing.T) {
	tests := []struct {
		para []int
		ret  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{2, 7, 9, 3, 1}, 22},
		{[]int{}, 0},
		{[]int{-1, -2}, -1},
	}
	for _, test := range tests {
		ret := maxSubArray(test.para)
		t.Logf("input: %v, output: %v", test.para, ret)

		if ret != test.ret {
			t.Errorf("expectd %v, actual %v", test.ret, ret)
		}
	}
}
