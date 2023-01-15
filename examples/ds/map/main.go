package main

import (
	"fmt"
	"log"

	ds "github.com/sonirico/stadio/ds/map"
)

func main() {
	var native ds.Map[string, int]

	native = ds.NewConcurrent[string, int](ds.NewNative[string, int]())
	native.Set("transit", 0)
	native.Set("huahuita", 1)
	native.Set("cdcompilador", 2)
	other := native.Map(func(s string, t int) (string, int) {
		return s + fmt.Sprintf("%d", t), t + 3
	})
	m := ds.NewConcurrent[string, int](other)

	m.Range(func(s string, t int, i int) bool {
		log.Println(s, t, i)
		return true
	})

}
