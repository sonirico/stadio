// Package maps provides utilities to work with maps
package maps

import (
	"github.com/sonirico/stadio/fp"
	"github.com/sonirico/stadio/slices"
	"github.com/sonirico/stadio/tuples"
)

// Equals returns whether 2 maps are equals in values
func Equals[K comparable, V any](m1, m2 map[K]V, eq func(V, V) bool) bool {
	if len(m1) != len(m2) {
		return false
	}

	if m1 == nil && m2 != nil {
		return false
	}

	if m1 != nil && m2 == nil {
		return false
	}

	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok {
			return false
		}

		if !eq(v1, v2) {
			return false
		}
	}

	return true
}

// Map transforms a map into another one, with same or different types
func Map[K1 comparable, V1 any, K2 comparable, V2 any](
	m map[K1]V1,
	p func(K1, V1) (K2, V2),
) map[K2]V2 {
	if m == nil {
		return nil
	}

	res := make(map[K2]V2, len(m))

	for k1, v1 := range m {
		k2, v2 := p(k1, v1)
		res[k2] = v2
	}

	return res
}

// FilterMap both filters and maps a map. The predicate function should return a fp.Option monad:
// fp.Some to indicate the entry should be kept.
// fp.None to indicate the entry should be discarded
func FilterMap[K1 comparable, V1 any, K2 comparable, V2 any](
	m map[K1]V1,
	p func(K1, V1) fp.Option[tuples.Tuple2[K2, V2]],
) map[K2]V2 {
	if m == nil {
		return nil
	}

	res := make(map[K2]V2, len(m))

	for k1, v1 := range m {
		tpl := p(k1, v1)
		if tpl.IsSome() {
			v := tpl.UnwrapUnsafe()
			res[v.V1] = v.V2
		}
	}

	return res
}

// FilterMapTuple both filters and maps the given map by receiving a predicate
// which returns mapped values, and a boolean to indicate whether that entry
// should be kept.
func FilterMapTuple[K1 comparable, V1 any, K2 comparable, V2 any](
	m map[K1]V1,
	p func(K1, V1) (K2, V2, bool),
) map[K2]V2 {
	if m == nil {
		return nil
	}

	res := make(map[K2]V2, len(m))

	for k1, v1 := range m {
		if k2, v2, ok := p(k1, v1); ok {
			res[k2] = v2
		}
	}

	return res
}

// Filter discards those entries from the map that do not match predicate.
func Filter[K comparable, V any](
	m map[K]V,
	p func(K, V) bool,
) map[K]V {
	if m == nil {
		return nil
	}

	res := make(map[K]V, len(m))

	for k, v := range m {
		if p(k, v) {
			res[k] = v
		}
	}

	return res
}

// FilterInPlace deletes those entries from the map that do not match predicate.
func FilterInPlace[K comparable, V any](
	m map[K]V,
	p func(K, V) bool,
) map[K]V {
	if m == nil {
		return nil
	}

	for k, v := range m {
		if !p(k, v) {
			delete(m, k)
		}
	}

	return m
}

// Reduce compacts the given map into a single type
func Reduce[K comparable, V any, R any](
	m map[K]V,
	p func(R, K, V) R,
) R {
	var r R

	if m == nil {
		return r
	}

	for k, v := range m {
		r = p(r, k, v)
	}

	return r
}

// Fold compacts the given map into a single type by taking into account the initial value
func Fold[K comparable, V any, R any](
	m map[K]V,
	p func(R, K, V) R,
	initial R,
) R {
	if m == nil {
		return initial
	}

	r := initial

	for k, v := range m {
		r = p(r, k, v)
	}

	return r
}

// Slice converts a map into a slice
func Slice[K comparable, V, R any](
	m map[K]V,
	p func(K, V) R,
) slices.Slice[R] {
	res := make([]R, len(m))
	i := 0

	for k, v := range m {
		res[i] = p(k, v)
		i++
	}

	return res
}
