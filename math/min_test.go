package math

import (
	"testing"
)

func TestMin(t *testing.T) {
	a, b := 1, 2
	if Min(a, b) != 1 {
		t.Error(Min(a, b))
	}
}

func TestMax(t *testing.T) {
	a, b := 1, 2
	if Max(a, b) != 2 {
		t.Error(Max(a, b))
	}
}
