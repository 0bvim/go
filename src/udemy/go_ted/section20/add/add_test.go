package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(1, 2)
	if total != 3 {
		t.Errorf("Add(1,2)=%d; want 3", total)
	}
}

func TestSubtract(t *testing.T) {
	total := Subtract(1, 2)
	if total != -1 {
		t.Errorf("Subtract(1,2)=%d; want -1", total)

	}
}

func TestDoMathSub(t *testing.T) {
	total := DoMath(1, 2, Subtract)
	if total != -1 {
		t.Errorf("Subtract(1,2)=%d; want -1", total)
	}
}

func TestDoMathAdd(t *testing.T) {
	total := DoMath(1, 2, Add)
	if total != 3 {
		t.Errorf("Add(1, 2)=%d; want 3", total)
	}
}

func TestParadise(t *testing.T) {
	location := "Canada"
	result := Paradise(location)

	if result != "My idea of paradise is: "+location {
		t.Errorf("Paradise(%q)=%q; want %q", location, result, location)
	}
}
