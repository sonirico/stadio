package stadio

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
	Wrapped[T any] struct {
		value T
		idx   int
	}
)

func ToMapIdx[V any, K comparable](arr []V, predicate func(x V) K) map[K]Wrapped[V] {
	res := make(map[K]Wrapped[V], len(arr))

	for i, x := range arr {
		res[predicate(x)] = Wrapped[V]{value: x, idx: i}
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

func Filter[T any](arr []T, predicate func(t T) bool) []T {
	res := make([]T, 0, len(arr))

	for _, x := range arr {
		if predicate(x) {
			res = append(res, x)
		}
	}

	return res
}

func FilterMap[T, U any](arr []T, predicate func(t T) Option[U]) []U {
	res := make([]U, 0, len(arr))

	for _, x := range arr {
		o := predicate(x)
		if o.IsSome() {
			res = append(res, o.Unwrap())
		}
	}

	return res
}

func FilterInPlace[T any](arr []T, predicate func(t T) bool) []T {
	n := 0
	for i, x := range arr {
		if predicate(x) {
			if n != i {
				arr[i] = x
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
				arr[i] = x
			}
			n++
		}
	}

	arr = arr[:n]

	res := make([]T, n)

	copy(res, arr[:n])

	return res
}
