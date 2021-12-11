package advent08

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
			wantPart1: 26,
			wantPart2: 61229,
		},
		{
			name:      "input",
			wantPart1: 294,
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

func TestDecodeMap(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want map[string]string
	}{
		{
			name: "Example",
			str:  "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
			want: map[string]string{
				"d": "a",
				"e": "b",
				"a": "c",
				"f": "d",
				"g": "e",
				"b": "f",
				"c": "g",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line := parseLine(tt.str)
			got := line.decodeMap()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNonMatches(t *testing.T) {
	type args struct {
		str       string
		possibles string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "from complete",
			args: args{
				str: "abcdefg",
				possibles: "cf",
			},
			want: "cf",
		},
		{
			name: "out of order",
			args: args{
				str: "abcdefg",
				possibles: "fc",
			},
			want: "cf",
		},
		{
			name: "incomplete",
			args: args{
				str: "abcde",
				possibles: "cf",
			},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNonMatches(tt.args.str, tt.args.possibles); got != tt.want {
				t.Errorf("removeNonMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}