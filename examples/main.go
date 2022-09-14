package main

import (
	"log"
	"strconv"

	std "github.com/sonirico/stadio"
)

func main() {

	slice := std.Slice[int]([]int{1, 2, 3, 4})

	res := std.FilterMap(slice, func(x int) std.Option[string] {
		if x%2 == 1 {
			return std.None[string]()
		}
		return std.Some(strconv.FormatInt(int64(x), 10))
	})

	std.Slice[string](res).Range(func(x string, i int) bool {
		log.Println(i, x)
		return true
	})

}
