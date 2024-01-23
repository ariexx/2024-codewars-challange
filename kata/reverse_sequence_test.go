package kata

import "testing"

func TestReverseSeq(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "1",
			n:    1,
			want: []int{1},
		},
		{
			name: "10",
			n:    10,
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSeq(tt.n); !equal(got, tt.want) {
				t.Errorf("ReverseSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
