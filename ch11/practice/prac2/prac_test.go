package prac2

import (
	"reflect"
	"testing"
)

// 不知道咋写  就这样吧，，，
func TestInsert(t *testing.T) {
	var x, y IntSet
	stx := make(map[uint64]struct{})
	sty := make(map[uint64]struct{})
	y.Add(9)
	y.Add(42)
	sty[9] = struct{}{}
	sty[42] = struct{}{}

	x.Add(1)
	stx[1] = struct{}{}
	x.Add(144)
	stx[144] = struct{}{}
	if !reflect.DeepEqual(x.GetWords(), stx) {
		t.Errorf("%v = %v", x.GetWords(), stx)
	}
	okx := x.Has(2)
	_, ok := stx[2]
	if okx != ok {
		t.Errorf("%v = %v", ok, okx)
	}
}
