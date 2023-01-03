package max_square_in_gistogram

import "testing"

func Test_getMaxSquare(t *testing.T) {
	tests := []struct {
		name      string
		gistogram []int
		want      int
	}{
		{name: "#1", gistogram: []int{2, 7, 6, 9, 7, 5, 7, 3, 5}, want: 30},
		{name: "#2", gistogram: []int{7, 2, 1, 4, 5, 1, 3, 3}, want: 8},
		{name: "#3", gistogram: []int{2, 1, 4, 5, 1, 3, 3}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxSquare(tt.gistogram); got != tt.want {
				t.Errorf("getMaxSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
