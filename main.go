package main

import (
	"fmt"
	"github.com/jakecoffman/generics/math"
	"github.com/jakecoffman/generics/slices"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	log.SetFlags(log.Lshortfile)

	a, b := 1, 2
	log.Println("min 1, 2 is ", math.Min(a, b))
	c, d := 1.1, 2.2
	log.Println("min 1.1, 2.2 is", math.Min(c, d))

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	evens := slices.Filter(arr, func(i int) bool {
		return arr[i]%2 == 0
	})
	log.Println("Evens:", evens)

	strings := slices.Map(arr, func(i int) string {
		return "v-" + strconv.Itoa(arr[i])
	})
	log.Println("Strings:", strings)

	// MapWithPool is an easy way to rate-limit http calls
	slices.MapWithPool(arr, 4, func(i int) error {
		r, err := http.Get(fmt.Sprint("https://example.com?q=", arr[i]))
		if err != nil {
			log.Println(err)
			return err
		}
		io.Copy(io.Discard, r.Body)
		return nil
	})

	things := []Thing{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "Daisy"},
	}
	i := slices.IndexOf(things, func(i int) bool {
		return things[i].Name == "Charlie"
	})
	log.Println("Found index:", i)
	element := slices.Find(things, func(i int) bool {
		return things[i].Name == "Charlie"
	})
	log.Printf("Found: %#v\n", element)
}

type Thing struct {
	ID   int
	Name string
}
