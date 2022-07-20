package consecutive

import (
	"fmt"
	"testing"
)

func TestRegression(t *testing.T) {
	tests := []struct {
		inout []int
	}{
		{[]int{5, 2}},
		{[]int{1, 1}},
		{[]int{3, 2}},
		{[]int{6, 2}},
		{[]int{2, 1}},
		{[]int{9, 3}},
		{[]int{15, 4}},
		{[]int{13, 2}},
		{[]int{22, 2}}, // 22= 4+5+6+7
		{[]int{23, 2}},
		{[]int{24, 2}},
		{[]int{99, 6}},
		{[]int{100, 3}},
		{[]int{11000023, 4}},
		{[]int{933320757, 16}},
	}

	for _, test := range tests {

		t.Run(fmt.Sprintf("Case %d expected %d", test.inout[0], test.inout[1]), func(t *testing.T) {
			got := consecutiveNumbersSum(test.inout[0])
			want := test.inout[1]
			if want != got {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
