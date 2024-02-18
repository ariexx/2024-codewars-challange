package kata

import "testing"

func Test_Rps(t *testing.T) {
	tests := []struct {
		p1, p2 string
		want   string
	}{
		{"rock", "scissors", "Player 1 won!"},
		{"scissors", "paper", "Player 1 won!"},
		{"paper", "rock", "Player 1 won!"},
		{"rock", "rock", "Draw!"},
		{"scissors", "scissors", "Draw!"},
	}

	for _, tt := range tests {
		t.Run(tt.p1+" vs "+tt.p2, func(t *testing.T) {
			if got := Rps(tt.p1, tt.p2); got != tt.want {
				t.Errorf("Rps(%q, %q) = %q; want %q", tt.p1, tt.p2, got, tt.want)
			}
		})
	}
}
