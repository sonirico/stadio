package _map

import (
	"github.com/sonirico/stadio/fp"
	"github.com/sonirico/stadio/maps"
	"github.com/sonirico/stadio/slices"
	"github.com/sonirico/stadio/tuples"
)

type (
	Entry[K comparable, V any] struct {
		K K
		V V
	}

	Native[K comparable, V any] struct {
		data map[K]V
	}
)

func NewNative[K comparable, V any]() Native[K, V] {
	return Native[K, V]{data: make(map[K]V)}
}

func (m Native[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m.data[k]
	return
}

func (m Native[K, V]) Has(k K) (ok bool) {
	_, ok = m.data[k]
	return
}

func (m Native[K, V]) Set(k K, v V) {
	m.data[k] = v
	return
}

func (m Native[K, V]) Range(fn func(K, V, int) bool) {
	i := 0
	for k, v := range m.data {
		if !fn(k, v, i) {
			return
		}
		i++
	}
}

func (m Native[K, V]) Delete(k K) {
	delete(m.data, k)
}

func (m Native[K, V]) GetOrSet(k K, def V) (v V, ok bool) {
	if v, ok = m.data[k]; ok {
		return
	}

	m.data[k] = def
	v = def
	ok = true
	return
}

func (m Native[K, V]) Map(fn func(K, V) (K, V)) Map[K, V] {
	return Native[K, V]{data: maps.Map(m.data, fn)}
}

func (m Native[K, V]) FilterMap(fn func(K, V) fp.Option[tuples.Tuple2[K, V]]) Map[K, V] {
	return Native[K, V]{data: maps.FilterMap(m.data, fn)}
}

func (m Native[K, V]) Filter(fn func(K, V) bool) Map[K, V] {
	return Native[K, V]{data: maps.Filter(m.data, fn)}
}

func (m Native[K, V]) Values() slices.Slice[V] {
	res := make([]V, len(m.data))
	i := 0
	for _, v := range m.data {
		res[i] = v
		i++
	}
	return res
}

func (m Native[K, V]) Keys() slices.Slice[K] {
	res := make([]K, len(m.data))
	i := 0
	for k := range m.data {
		res[i] = k
		i++
	}
	return res
}

func (m Native[K, V]) Entries() slices.Slice[Entry[K, V]] {
	res := make([]Entry[K, V], len(m.data))
	i := 0
	for k, v := range m.data {
		res[i] = Entry[K, V]{K: k, V: v}
		i++
	}
	return res
}
