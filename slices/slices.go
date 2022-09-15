package slices

import (
	"github.com/sonirico/stadio/fp"
)

type (
	Slice[T any] []T
)

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Range(fn func(t T, i int) bool) {
	for i, x := range s {
		if !fn(x, i) {
			return
		}
	}
}

func (s Slice[T]) Get(i int) (res T, ok bool) {
	ok = i >= 0 && i < len(s)
	if !ok {
		return
	}
	res = s[i]
	return
}

func (s Slice[T]) Contains(fn func(t T) bool) bool {
	return Contains(s, fn)
}

func (s Slice[T]) Equals(other Slice[T], predicate func(x, y T) bool) (res bool) {
	return Equals(s, other, predicate)
}

func (s Slice[T]) Clone() Slice[T] {
	res := make([]T, len(s))
	copy(res, s)
	return res
}

func (s Slice[T]) Map(predicate func(T) T) Slice[T] {
	return Map(s, predicate)
}

func (s Slice[T]) MapInPlace(predicate func(T) T) Slice[T] {
	return MapInPlace(s, predicate)
}

func (s Slice[T]) Filter(predicate func(x T) bool) Slice[T] {
	return Filter(s, predicate)
}

func (s Slice[T]) FilterMapTuple(predicate func(x T) (T, bool)) Slice[T] {
	return FilterMapTuple(s, predicate)
}

func (s Slice[T]) FilterMap(predicate func(x T) fp.Option[T]) Slice[T] {
	return FilterMap(s, predicate)
}

func (s Slice[T]) FilterInPlace(predicate func(x T) bool) Slice[T] {
	return FilterInPlace(s, predicate)
}

func (s Slice[T]) FilterInPlaceCopy(predicate func(x T) bool) Slice[T] {
	return FilterInPlaceCopy(s, predicate)
}

func (s Slice[T]) Reduce(predicate func(x, y T) T) T {
	return ReduceSame(s, predicate)
}

func (s Slice[T]) Fold(predicate func(x, y T) T, initial T) T {
	return FoldSame(s, predicate, initial)
}

func Equals[T any](one, other []T, predicate func(x, y T) bool) (res bool) {
	if len(one) != len(other) {
		return
	}

	res = true

	for idx, otherItem := range other {
		res = predicate(one[idx], otherItem)
		if !res {
			return
		}
	}
	return
}

func (s Slice[T]) IndexOf(fn func(t T) bool) int {
	return IndexOf(s, fn)
}

func ToMap[V any, K comparable](arr []V, predicate func(x V) K) map[K]V {
	res := make(map[K]V, len(arr))

	for _, x := range arr {
		res[predicate(x)] = x
	}

	return res
}

type (
	WrappedIdx[T any] struct {
		value T
		idx   int
	}
)

func ToMapIdx[V any, K comparable](arr []V, predicate func(x V) K) map[K]WrappedIdx[V] {
	res := make(map[K]WrappedIdx[V], len(arr))

	for i, x := range arr {
		res[predicate(x)] = WrappedIdx[V]{value: x, idx: i}
	}

	return res
}

func IndexOf[T any](arr []T, predicate func(t T) bool) (pos int) {
	pos = -1
	for i, x := range arr {
		if predicate(x) {
			pos = i
			return
		}
	}
	return
}

func Contains[T any](arr []T, predicate func(t T) bool) bool {
	return IndexOf(arr, predicate) >= 0
}

func Map[T, U any](arr []T, predicate func(t T) U) []U {
	res := make([]U, 0, len(arr))

	for _, x := range arr {
		res = append(res, predicate(x))
	}

	return res
}

func MapInPlace[T any](arr []T, predicate func(t T) T) []T {
	for i, x := range arr {
		arr[i] = predicate(x)
	}

	return arr
}

func Filter[T any](arr []T, predicate func(t T) bool) []T {
	res := make([]T, 0, len(arr))

	for _, x := range arr {
		if predicate(x) {
			res = append(res, x)
		}
	}

	return res
}

func FilterMapTuple[T, U any](arr []T, predicate func(t T) (U, bool)) []U {
	res := make([]U, 0, len(arr))

	for _, x := range arr {
		if mapped, ok := predicate(x); ok {
			res = append(res, mapped)
		}
	}

	return res
}

func FilterMap[T, U any](arr []T, predicate func(t T) fp.Option[U]) []U {
	pre := func(t T) (U, bool) {
		return predicate(t).Unwrap()
	}

	return FilterMapTuple[T, U](arr, pre)
}

func FilterInPlace[T any](arr []T, predicate func(t T) bool) []T {
	n := 0
	for i, x := range arr {
		if predicate(x) {
			if n != i {
				arr[n] = x
			}
			n++
		}
	}

	arr = arr[:n]

	return arr
}

func FilterInPlaceCopy[T any](arr []T, predicate func(t T) bool) []T {
	n := 0
	for i, x := range arr {
		if predicate(x) {
			if n != i {
				arr[n] = x
			}
			n++
		}
	}

	arr = arr[:n]

	res := make([]T, n)

	copy(res, arr[:n])

	return res
}

func Reduce[T, U any](arr []T, p func(T, T) T) (res T) {
	return Fold(arr, p, res)
}

func ReduceSame[T any](arr []T, p func(T, T) T) T {
	return Reduce[T, T](arr, p)
}

func FoldSame[T any](arr []T, p func(T, T) T, initial T) T {
	return Fold[T, T](arr, p, initial)
}

func Fold[T, U any](arr []T, p func(U, T) U, initial U) U {
	if len(arr) < 1 {
		return initial
	}

	initial = p(initial, arr[0])

	if len(arr) < 2 {
		return initial
	}

	i := 1

	for i < len(arr) {
		initial = p(initial, arr[i])

		i++
	}

	return initial
}

// Cut removes a sector from slice given lower and upper bounds. Bounds are
// represented as indices of the slice. E.g:
// Cut([1, 2, 3, 4], 1, 2) -> [1, 4]
// Cut([4], 0, 0) -> []
// Cut will returned the original slice without the cut subslice.
func Cut[T any](arr []T, from, to int) []T {
	if len(arr) < 1 {
		return arr
	}

	if from < 0 {
		from = 0
	}

	if from >= len(arr) {
		from = len(arr) - 1
	}

	if to < 0 {
		to = 0
	}

	if to >= len(arr) {
		to = len(arr) - 1
	}

	if len(arr) == 1 {
		return arr[:0]
	}

	if from > to {
		// In this case, consider `to` to be the amount to remove from `from`.
		return append(arr[:from], arr[from+to+1:]...)
	}

	return append(arr[:from], arr[to+1:]...)
}
