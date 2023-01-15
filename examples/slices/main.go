package main

import (
	"log"
	"strconv"

	std "github.com/sonirico/stadio/fp"
	"github.com/sonirico/stadio/slices"
)

func main() {

	slice := slices.Slice[int]([]int{1, 2, 3, 4})

	res := slices.FilterMap(slice, func(x int) std.Option[string] {
		if x%2 == 1 {
			return std.None[string]()
		}
		return std.Some(strconv.FormatInt(int64(x), 10))
	})

	slices.Slice[string](res).Range(func(x string, i int) bool {
		log.Println(i, x)
		return true
	})

}
