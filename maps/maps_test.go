package maps

import (
	"reflect"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	v := Keys(m)
	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	if !reflect.DeepEqual(v, []string{"a", "b", "c"}) {
		t.Error(v)
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	v := Values(m)
	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	if !reflect.DeepEqual(v, []int{1, 2, 3}) {
		t.Error(v)
	}
}

func TestAssign(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
	}
	m2 := map[string]int{
		"b": 2,
	}
	v := Assign(map[string]int{}, m1, m2)
	if !reflect.DeepEqual(v, map[string]int{
		"a": 1,
		"b": 2,
	}) {
		t.Error(v)
	}
}
