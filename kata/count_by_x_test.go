package kata

import "testing"

func TestCountBy(t *testing.T) {
	tests := []struct {
		x, n   int
		result []int
		want   []int
	}{
		{1, 5, []int{}, []int{1, 2, 3, 4, 5}},
		{2, 5, []int{}, []int{2, 4, 6, 8, 10}},
		{3, 5, []int{}, []int{3, 6, 9, 12, 15}},
	}

	for _, test := range tests {
		if got := CountBy(test.x, test.n); !equal(got, test.want) {
			t.Errorf("CountBy(%d, %d) = %v, want %v", test.x, test.n, got, test.want)
		}
	}
}
