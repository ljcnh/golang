package prac1

import (
	"reflect"
	"testing"
)

func TestCharcount(t *testing.T) {
	t1 := map[rune]int{
		102: 2, 10: 1, 113: 1, 103: 2, 118: 1, 97: 2, 115: 2, 119: 2, 51: 1, 100: 3, 101: 2,
	}
	t2 := map[rune]int{
		97: 2,
	}
	t3 := map[rune]int{
		107: 2, 97: 2, 121: 1,
	}
	t4 := map[rune]int{
		233: 2, 116: 1,
	}
	t5 := map[rune]int{
		114: 5, 105: 2, 46: 1, 116: 2, 32: 4, 101: 7, 118: 2, 44: 1, 69: 1, 115: 6,
	}
	var tests = []struct {
		input string
		want  map[rune]int
	}{
		{"asddf\nsadew3qfgwegv", t1},
		{"aa", t2},
		{"kayak", t3},
		{"été", t4},
		{"Et se resservir, ivresse reste.", t5},
	}
	for _, test := range tests {
		if got := Charcount(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}

}
