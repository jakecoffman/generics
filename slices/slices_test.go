package slices

import (
	"reflect"
	"runtime"
	"testing"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	arr = Filter(arr, func(i int) bool {
		return arr[i]%2 == 0
	})
	if !reflect.DeepEqual(arr, []int{2, 4}) {
		t.Error(arr)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	arr2 := Map(arr, func(i int) bool {
		return arr[i]%2 == 0
	})
	if !reflect.DeepEqual(arr2, []bool{false, true, false, true}) {
		t.Error(arr)
	}
}

func TestMapP(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	arr2 := MapWithPool(arr, 2, func(i int) bool {
		return arr[i]%2 == 0
	})
	if !reflect.DeepEqual(arr2, []bool{false, true, false, true}) {
		t.Error(arr)
	}
}

func BenchmarkMap(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	expensiveOperation := func(i int) bool {
		v := arr[i]
		for j := 0; j < 10000; j++ {
			v++
		}
		return v%2 == 0
	}

	b.Run("Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Map(arr, expensiveOperation)
		}
	})
	b.Run("MapWithPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MapWithPool(arr, runtime.NumCPU(), expensiveOperation)
		}
	})
}

func TestSome(t *testing.T) {
	arr := []int{1, 3}
	if v := Some(arr, func(i int) bool {
		return arr[i]%2 == 0
	}); v {
		t.Error()
	}
	if v := Some(arr, func(i int) bool {
		return arr[i]%2 != 0
	}); !v {
		t.Error()
	}
}

func TestReduce(t *testing.T) {
	var v []int
	add := func(a, b int) int {
		return a + b
	}
	if r := Reduce(v, add); r != 0 {
		t.Error(r)
	}
	v = []int{1}
	if r := Reduce(v, add); r != 1 {
		t.Error(r)
	}
	v = []int{1, 9, 8, 2, 7, 3, 6, 4, 5, 5}
	if r := Reduce(v, add); r != 50 {
		t.Error(r)
	}
}

func TestUnique(t *testing.T) {
	v := []int{1, 1, 2, 3, 3, 3, 3, 4}
	if r := Unique(v); !reflect.DeepEqual(r, []int{1, 2, 3, 4}) {
		t.Error(r)
	}
	v = nil
	if r := Unique(v); !reflect.DeepEqual(r, []int{}) {
		t.Error(r)
	}
}

func TestForEach(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7}
	ForEach(v, func(e *int) {
		*e += 10
	})
	if !reflect.DeepEqual(v, []int{11, 12, 13, 14, 15, 16, 17}) {
		t.Error(v)
	}
}

func TestForEachWithPool(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7}
	ForEachWithPool(v, 4, func(e *int) {
		*e += 10
	})
	if !reflect.DeepEqual(v, []int{11, 12, 13, 14, 15, 16, 17}) {
		t.Error(v)
	}
}

func TestPop(t *testing.T) {
	v := []int{1, 2}
	two := Pop(&v)
	if *two != 2 {
		t.Error(two)
	}
	if !reflect.DeepEqual(v, []int{1}) {
		t.Error(v)
	}
	one := Pop(&v)
	if *one != 1 {
		t.Error(one)
	}
	if !reflect.DeepEqual(v, []int{}) {
		t.Error(v)
	}
	if Pop(&v) != nil {
		t.Error()
	}
	if !reflect.DeepEqual(v, []int{}) {
		t.Error(v)
	}
}

func TestPush(t *testing.T) {
	var v []int
	Push(&v, 1, 2, 3, 4)
	if !reflect.DeepEqual(v, []int{1, 2, 3, 4}) {
		t.Error(v)
	}
	Push(&v, 5, 6)
	if !reflect.DeepEqual(v, []int{1, 2, 3, 4, 5, 6}) {
		t.Error(v)
	}
}

func TestUnshift(t *testing.T) {
	var v []int
	Unshift(&v, 1, 2, 3, 4)
	if !reflect.DeepEqual(v, []int{1, 2, 3, 4}) {
		t.Error(v)
	}
	Unshift(&v, 5, 6)
	if !reflect.DeepEqual(v, []int{5, 6, 1, 2, 3, 4}) {
		t.Error(v)
	}
}

func TestShift(t *testing.T) {
	v := []int{1, 2}
	Shift(&v)
	if !reflect.DeepEqual(v, []int{2}) {
		t.Error(v)
	}
	Shift(&v)
	if !reflect.DeepEqual(v, []int{}) {
		t.Error(v)
	}
	Shift(&v)
	if !reflect.DeepEqual(v, []int{}) {
		t.Error(v)
	}
}

func TestSplice(t *testing.T) {
	v := []int{0, 1, 2, 3, 4}
	removed := Splice(&v, 2, 2, 9, 9, 8)
	if !reflect.DeepEqual(v, []int{0, 1, 9, 9, 8, 4}) {
		t.Error(v)
	}
	if !reflect.DeepEqual(removed, []int{2, 3}) {
		t.Error(removed)
	}
}

func TestReverse(t *testing.T) {
	v := []int{0, 1, 2, 3}
	r := Reverse(v)
	if !reflect.DeepEqual(r, []int{3, 2, 1, 0}) {
		t.Error(v)
	}
	v = []int{0, 1, 2, 3, 4}
	r = Reverse(v)
	if !reflect.DeepEqual(r, []int{4, 3, 2, 1, 0}) {
		t.Error(v)
	}
}

// Performance of looping in this library is always going to be worse than
// writing the for loop because at the very least we're introducing a function
// call to each element of the array.
func BenchmarkForEach(b *testing.B) {
	const size = 1000
	v := make([]int, size)
	for i := 0; i < size; i++ {
		v[i] = i
	}

	b.Run("ForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ForEach(v, func(e *int) {
				*e += 1
			})
		}
	})

	// reset in case it matters
	for i := 0; i < size; i++ {
		v[i] = i
	}

	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for index := range v {
				// I tried keeping this as close to the implementation of ForEach as above
				e := &v[index]
				*e += 1
			}
		}
	})
}
