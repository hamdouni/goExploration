package sum_test

import "testing"
import "goExploration/sum"

func TestSum(t *testing.T) {
	tt := []struct {
		val []int
		res int
		lab string
	}{
		{[]int{}, 0, "sum of none = 0 "},
		{[]int{1, 2, 3, 4, 5, 6}, 21, "sum 1 to 6 = 21"},
		{[]int{1, -1}, 0, "sum 1 -1 = 0"},
		{[]int{4, 5, 6}, 15, "sum 4 to 6 = 15"},
	}
	for _, tc := range tt {
		s := sum.Sum(tc.val...)
		if s != tc.res {
			t.Errorf("Expect %v got %v", tc.lab, s)
		}
	}
}
