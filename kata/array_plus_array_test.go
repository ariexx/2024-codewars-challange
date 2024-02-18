package kata

import "testing"

func TestArrayPArray(t *testing.T) {
	tests := []struct {
		arr1, arr2 []int
		want       int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, 21},
		{[]int{0, 0, 0}, []int{4, 5, 6}, 15},
		{[]int{1, 2, 3}, []int{0, 0, 0}, 6},
		{[]int{0, 0, 0}, []int{0, 0, 0}, 0},
		{[]int{16, 17, 18}, []int{19, 20, 21}, 111},
	}

	for _, tt := range tests {
		t.Run("ArrayPArray", func(t *testing.T) {
			if got := ArrayPArray(tt.arr1, tt.arr2); got != tt.want {
				t.Errorf("ArrayPArray(%v, %v) = %v; want %v", tt.arr1, tt.arr2, got, tt.want)
			}
		})
	}
}
