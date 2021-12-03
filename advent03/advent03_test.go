package advent03

import (
	"reflect"
	"strconv"
	"testing"

	"advent2021/util"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		wantPart1 interface{}
		wantPart2 interface{}
	}{
		{
			name: "simple",
			wantPart1: 198,
			wantPart2: 230,
		},
		{
			name: "input",
			wantPart1: 4174964,
			wantPart2: 4474944,
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

func TestBinary(t *testing.T) {
	 r := binary("10110")
	 if r != 22 {
		 t.Errorf("incorrect %d", r)
	 }

	r = binary("01001")
	if r != 9 {
		t.Errorf("incorrect %d", r)
	}

	r = uint8(22)
	r = r ^ 0x1F
	if r != 9 {
		t.Errorf("incorrect %d", r)
	}
}

func binary(str string) uint8 {
	v, err := strconv.ParseUint(str, 2, 5)
	util.Panic(err)
	return uint8(v)
}
