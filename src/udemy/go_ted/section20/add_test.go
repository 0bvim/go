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
