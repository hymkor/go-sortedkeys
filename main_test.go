package sortedkeys

import (
	"testing"
)

func TestRange(t *testing.T) {
	pattern := map[string]int{
		"B": 1,
		"A": 0,
		"D": 3,
		"C": 2,
	}
	p := New(pattern)
	if !p.Range() || p.Key != "A" || p.Value != 0 {
		t.Fail()
	}
	if !p.Range() || p.Key != "B" || p.Value != 1 {
		t.Fail()
	}
	if !p.Range() || p.Key != "C" || p.Value != 2 {
		t.Fail()
	}
	if !p.Range() || p.Key != "D" || p.Value != 3 {
		t.Fail()
	}
	if p.Range() {
		t.Fail()
	}
}
