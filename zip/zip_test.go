package zip

import "testing"

func TestZip_Next(t *testing.T) {
	z := New([]int{1, 2, 3}, []string{"one", "two", "three"})
	for z.Next() {
		t.Log(z.Value())
	}
}
