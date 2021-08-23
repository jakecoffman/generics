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
