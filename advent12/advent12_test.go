package advent12

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		wantPart1 interface{}
		wantPart2 interface{}
	}{
		{
			name: "simple",
			wantPart1: 10,
			wantPart2: 0,
		},
		{
			name: "simple2",
			wantPart1: 19,
			wantPart2: 0,
		},
		{
			name: "simple3",
			wantPart1: 226,
			wantPart2: 0,
		},
		{
			name: "input",
			wantPart1: 3713,
			wantPart2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPart1, gotPart2 := Solution(tt.name + ".txt")
			if !reflect.DeepEqual(gotPart1, tt.wantPart1) {
				t.Errorf("Solution() gotPart1 = %v, want %v", gotPart1, tt.wantPart1)
			}
			if !reflect.DeepEqual(gotPart2, tt.wantPart2) {
				t.Errorf("Solution() gotPart2 = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}
