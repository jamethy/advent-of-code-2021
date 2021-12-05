package advent05

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
			name:      "simple",
			wantPart1: 5,
			wantPart2: 12,
		},
		{
			name:      "input",
			wantPart1: 5585,
			wantPart2: 17193,
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

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		line  Line
		point Point
		want  bool
	}{
		{
			name:  "left-right true",
			line:  Line{Start: Point{X: 1, Y: 2}, End: Point{X: 4, Y: 2}},
			point: Point{X: 2, Y: 2},
			want:  true,
		},
		{
			name:  "left-right before",
			line:  Line{Start: Point{X: 1, Y: 2}, End: Point{X: 4, Y: 2}},
			point: Point{X: 0, Y: 2},
			want:  false,
		},
		{
			name:  "left-right after",
			line:  Line{Start: Point{X: 1, Y: 2}, End: Point{X: 4, Y: 2}},
			point: Point{X: 8, Y: 2},
			want:  false,
		},
		{
			name:  "up-down true",
			line:  Line{Start: Point{X: 1, Y: 2}, End: Point{X: 1, Y: 5}},
			point: Point{X: 1, Y: 3},
			want:  true,
		},
		{
			name:  "up-down before",
			line:  Line{Start: Point{X: 1, Y: 2}, End: Point{X: 1, Y: 5}},
			point: Point{X: 1, Y: 1},
			want:  false,
		},
		{
			name:  "up-down after",
			line:  Line{Start: Point{X: 1, Y: 2}, End: Point{X: 1, Y: 5}},
			point: Point{X: 1, Y: 8},
			want:  false,
		},
		{
			name:  "lower-left-upper-right true",
			line:  Line{Start: Point{X: 1, Y: 1}, End: Point{X: 4, Y: 4}},
			point: Point{X: 2, Y: 2},
			want:  true,
		},
		{
			name:  "lower-left-upper-right start",
			line:  Line{Start: Point{X: 1, Y: 1}, End: Point{X: 4, Y: 4}},
			point: Point{X: 1, Y: 1},
			want:  true,
		},
		{
			name:  "lower-left-upper-right end",
			line:  Line{Start: Point{X: 1, Y: 1}, End: Point{X: 4, Y: 4}},
			point: Point{X: 4, Y: 4},
			want:  true,
		},
		{
			name:  "lower-left-upper-right before",
			line:  Line{Start: Point{X: 1, Y: 1}, End: Point{X: 4, Y: 4}},
			point: Point{X: 0, Y: 0},
			want:  false,
		},
		{
			name:  "lower-left-upper-right after",
			line:  Line{Start: Point{X: 1, Y: 1}, End: Point{X: 4, Y: 4}},
			point: Point{X: 5, Y: 5},
			want:  false,
		},
		{
			name:  "upper-right-lower-left true",
			line:  Line{Start: Point{X: 4, Y: 4}, End: Point{X: 1, Y: 1}},
			point: Point{X: 2, Y: 2},
			want:  true,
		},
		{
			name:  "upper-right-lower-left before",
			line:  Line{Start: Point{X: 4, Y: 4}, End: Point{X: 1, Y: 1}},
			point: Point{X: 5, Y: 5},
			want:  false,
		},
		{
			name:  "upper-right-lower-left after",
			line:  Line{Start: Point{X: 4, Y: 4}, End: Point{X: 1, Y: 1}},
			point: Point{X: 0, Y: 0},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.line.Intersects(tt.point.X, tt.point.Y)
			if got != tt.want {
				t.Errorf("Intersects() got = %v, want %v", got, tt.want)
			}
		})
	}
}
