package advent06

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		wantPart1 interface{}
		wantPart2 interface{}
	}{
		{
			name:      "simple",
			wantPart1: 5934,
			wantPart2: 26984457539,
		},
		{
			name:      "input",
			wantPart1: 372984,
			wantPart2: 1681503251694,
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

func TestCalcFish(t *testing.T) {
	tests := []struct {
		days int
		want int
	}{
		{
			days: 0,
			want: 0,
		},
		{
			days: 1,
			want: 0,
		},
		{
			days: 6,
			want: 0,
		},
		{
			days: 7,
			want: 1,
		},
		{
			days: 13,
			want: 1,
		},
		{
			days: 14,
			want: 2,
		},
		{
			days: 15,
			want: 2,
		},
		{
			days: 16,
			want: 3,
		},
		{
			days: 20,
			want: 3,
		},
		{
			days: 21,
			want: 4,
		},
		{
			days: 23,
			want: 6,
		},
		{
			days: 24,
			want: 6,
		},
		{
			days: 25,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.days), func(t *testing.T) {
			got := calcChildren(tt.days)
			if got != tt.want {
				t.Errorf("calcChildren got = %v, want %v", got, tt.want)
			}
		})
	}
}
