package main

import (
	"fmt"
	"github.com/jakecoffman/generics/math"
	"github.com/jakecoffman/generics/stream"
)

func main() {
	a, b := 1, 2
	if math.Min(a, b) != 1 {
		panic("Unexpected")
	}
	c, d := 1.1, 2.2
	if math.Min(c, d) != 1.1 {
		panic("Unexpected")
	}

	arr1 := []int{1, 2, 3, 4}
	v := stream.New(arr1).
		Filter(func(item *int) bool {
			return *item%2 != 0
		}).
		ToSlice()
	fmt.Println(v)

	arr2 := []string{"one", "two", "three", "four"}
	v2 := stream.New(arr2).
		Filter(func(item *string) bool {
			return len(*item) < 4
		}).
		ToSlice()
	fmt.Println(v2)

	type User struct {
		Name     string
		Internal bool
	}
	arr3 := []User{
		{Name: "Bob", Internal: true},
		{Name: "Alice"},
		{Name: "Cynthia"},
		{Name: "Derrick", Internal: true},
	}
	v3 := stream.New(arr3).
		Filter(func(item *User) bool {
			return item.Internal
		}).
		ToSlice()

	fmt.Println(v3)
}
