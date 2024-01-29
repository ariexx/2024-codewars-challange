package kata

import "testing"

func TestDNAtoRNA(t *testing.T) {
	test := []struct {
		name string
		dna  string
		want string
	}{
		{
			name: "1",
			dna:  "TTTT",
			want: "UUUU",
		},
		{
			name: "2",
			dna:  "GCAT",
			want: "GCAU",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if got := DNAtoRNA(tt.dna); got != tt.want {
				t.Errorf("DNAtoRNA() = %v, want %v", got, tt.want)
			}
		})
	}
}
