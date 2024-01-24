package kata

import (
	"testing"
)

func TestHero(t *testing.T) {

	tests := []struct {
		name   string
		dragon int
		bullet int
		want   bool
	}{
		{
			name:   "7 Dragon",
			dragon: 7,
			bullet: 10,
			want:   false,
		},
		{
			name:   "2 Dragon",
			dragon: 2,
			bullet: 4,
			want:   true,
		},
		{
			name:   "3 Dragon",
			dragon: 3,
			bullet: 1,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			if got = Hero(tt.bullet, tt.dragon); got != tt.want {
				t.Errorf("Hero fight %d dragons and carrying %d bullets, want result is %v got %v", tt.dragon, tt.bullet, tt.want, got)
			}
		})
	}
}
