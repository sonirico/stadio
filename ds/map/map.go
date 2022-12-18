package _map

import (
	"github.com/sonirico/stadio/fp"
	"github.com/sonirico/stadio/slices"
	"github.com/sonirico/stadio/tuples"
)

type (
	Map[K comparable, V any] interface {
		Get(K) (V, bool)
		Has(K) bool
		Set(K, V)
		Range(fn func(K, V, int) bool)
		Delete(K)
		GetOrSet(K, V) (V, bool)
		Map(fn func(K, V) (K, V)) Map[K, V]
		FilterMap(fn func(K, V) fp.Option[tuples.Tuple2[K, V]]) Map[K, V]
		Filter(fn func(K, V) bool) Map[K, V]
		Keys() slices.Slice[K]
		Values() slices.Slice[V]
		Entries() slices.Slice[Entry[K, V]]
	}
)
