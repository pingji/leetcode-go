package leetcode

import "testing"

func TestProblem(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	tests := []struct {
		i      int
		j      int
		output int
	}{
		{0, 2, 1},
		{2, 5, -1},
		{0, 5, -3},
	}
	for _, test := range tests {
		output := obj.SumRange(test.i, test.j)
		t.Logf("i: %v, j: %v, output %v, expected: %v", test.i, test.j, output, test.output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}
