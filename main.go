package main

import (
	"fmt"
	"github.com/jakecoffman/generics/math"
)

func main() {
	a, b := 1, 2
	if math.Min(a, b) != 1 {
		panic("Unexpected")
	}
	fmt.Println("OK!")
}
