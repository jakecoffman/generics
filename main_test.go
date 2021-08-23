package main

import (
	"github.com/jakecoffman/generics/math"
	"testing"
)

func TestMath(t *testing.T) {
	a, b := 1, 2
	if math.Min(a, b) != 1 {
		t.Error(math.Min(a, b))
	}
}
