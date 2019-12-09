package common_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/oizgagin/adventofcode2019/common"
)

func TestPermutations(t *testing.T) {

	testCases := []struct {
		n    int
		want [][]int
	}{
		{1, [][]int{[]int{0}}},
		{2, [][]int{[]int{0, 1}, []int{1, 0}}},
		{3, [][]int{
			[]int{0, 1, 2},
			[]int{0, 2, 1},
			[]int{2, 0, 1},
			[]int{1, 0, 2},
			[]int{1, 2, 0},
			[]int{2, 1, 0},
		}},
		{4, [][]int{
			[]int{0, 1, 2, 3},
			[]int{0, 1, 3, 2},
			[]int{0, 2, 1, 3},
			[]int{0, 2, 3, 1},
			[]int{0, 3, 1, 2},
			[]int{0, 3, 2, 1},
			[]int{1, 0, 2, 3},
			[]int{1, 0, 3, 2},
			[]int{1, 2, 0, 3},
			[]int{1, 2, 3, 0},
			[]int{1, 3, 0, 2},
			[]int{1, 3, 2, 0},
			[]int{2, 0, 1, 3},
			[]int{2, 0, 3, 1},
			[]int{2, 1, 0, 3},
			[]int{2, 1, 3, 0},
			[]int{2, 3, 0, 1},
			[]int{2, 3, 1, 0},
			[]int{3, 0, 1, 2},
			[]int{3, 0, 2, 1},
			[]int{3, 1, 0, 2},
			[]int{3, 1, 2, 0},
			[]int{3, 2, 0, 1},
			[]int{3, 2, 1, 0},
		}},
	}

	for _, tc := range testCases {
		got := common.Permutations(tc.n)

		sort.Slice(got, func(i, j int) bool { return lexless(got[i], got[j]) })
		sort.Slice(tc.want, func(i, j int) bool { return lexless(tc.want[i], tc.want[j]) })

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%d: got %v, want %v", tc.n, got, tc.want)
		}
	}

}

func lexless(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			continue
		}
		return a[i] < b[i]
	}
	return false
}
