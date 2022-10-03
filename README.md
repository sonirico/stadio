
# Stadio

Compendium of functions, data structures, monadic wrappers and more which, hopefully, will be
included as a standard library of the language

## Modules

### Slices

Package slices provides utilities to work with slices

Table of contents

- [Cut](####Cut)
- [Delete](####Delete)
- [DeleteOrder](####DeleteOrder)
- [Extract](####Extract)
- [ExtractIdx](####ExtractIdx)
- [Find](####Find)
- [FindIdx](####FindIdx)
- [Insert](####Insert)
- [InsertVector](####InsertVector)
- [Peek](####Peek)
- [Pop](####Pop)
- [PopFront](####PopFront)
- [PushFront](####PushFront)
- [Shift](####Shift)
- [Unshift](####Unshift)

#### Cut

Cut removes a sector from slice given lower and upper bounds. Bounds are
represented as indices of the slice. E.g:
Cut([1, 2, 3, 4], 1, 2) -> [1, 4]
Cut([4], 0, 0) -> []
Cut will returned the original slice without the cut subslice.


<details><summary>Code</summary>

```go

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

		return append(arr[:from], arr[from+to+1:]...)
	}

	return append(arr[:from], arr[to+1:]...)
}
```

</details>

#### Delete

Delete removes the element in `idx` position, without preserving array order. In case `idx`
is out of bounds, noop.


<details><summary>Code</summary>

```go

func Delete[T any](arr []T, idx int) []T {
	le := len(arr) - 1
	if le < 0 || idx > le || idx < 0 {
		return arr
	}
	var t T
	arr[idx] = arr[le]
	arr[le] = t
	arr = arr[:le]
	return arr
}
```

</details>

#### DeleteOrder

DeleteOrder removes the element in `idx` position, preserving array order. In case `idx`
is out of bounds, noop.


<details><summary>Code</summary>

```go

func DeleteOrder[T any](arr []T, idx int) []T {
	le := len(arr) - 1
	if le < 0 || idx > le || idx < 0 {
		return arr
	}
	var t T

	if le > 0 {
		copy(arr[idx:], arr[idx+1:])
	}

	arr[le] = t
	arr = arr[:le]
	return arr
}
```

</details>

#### Extract

Extract gets and deletes the element than matches predicate. Returned values are the
modified slice, the item or zero value if not found, and whether item was found


<details><summary>Code</summary>

```go

func Extract[T any](arr []T, predicate func(t T) bool) ([]T, T, bool) {
	res, idx := FindIdx(arr, predicate)
	if idx < 0 {
		return arr, res, false
	}

	arr = Delete(arr, idx)
	return arr, res, true
}
```

</details>

#### ExtractIdx

ExtractIdx gets and deletes the element at the given position. Returned values are the
modified slice, the item or zero value if not found, and whether item was found


<details><summary>Code</summary>

```go

func ExtractIdx[T any](arr []T, idx int) (res []T, item T, ok bool) {
	if idx >= len(arr) || idx < 0 {
		return
	}

	ok = true
	item = arr[idx]
	res = Delete(arr, idx)

	return
}
```

</details>

#### Find

Find returns the first element that matches predicate


<details><summary>Code</summary>

```go

func Find[T any](arr []T, predicate func(t T) bool) (res T, ok bool) {
	var idx int
	res, idx = FindIdx(arr, predicate)
	ok = idx > -1
	return
}
```

</details>

#### FindIdx

FindIdx returns the first element that matches predicate as well as the position on the slice.


<details><summary>Code</summary>

```go

func FindIdx[T any](arr []T, predicate func(t T) bool) (res T, idx int) {
	idx = IndexOf(arr, predicate)
	if idx < 0 {
		return
	}

	res = arr[idx]
	return
}
```

</details>

#### Insert

Insert places the given item at the position `idx` for the given slice


<details><summary>Code</summary>

```go

func Insert[T any](arr []T, item T, idx int) []T {
	if arr == nil {
		return []T{item}
	}

	if idx < 0 || idx > len(arr) {
		return arr
	}

	return append(arr[:idx], append([]T{item}, arr[idx:]...)...)
}
```

</details>

#### InsertVector

InsertVector places the given vector at the position `idx` for the given slice, moving
existing elements to the right.


<details><summary>Code</summary>

```go

func InsertVector[T any](arr, items []T, idx int) (res []T) {
	if arr == nil {
		res = items[:]
		return
	}

	if items == nil || len(items) == 0 {
		res = arr
		return
	}

	if idx < 0 || idx > len(arr) {
		return arr
	}

	return append(arr[:idx], append(items, arr[idx:]...)...)
}
```

</details>

#### Peek

Peek returns the item corresponding to idx


<details><summary>Code</summary>

```go

func Peek[T any](arr []T, idx int) (item T, ok bool) {
	if len(arr) < 1 || idx >= len(arr) {
		return
	}

	item = arr[idx]
	ok = true

	return
}
```

</details>

#### Pop

Pop deletes and returns the last item from the slice, starting from the end.


<details><summary>Code</summary>

```go

func Pop[T any](arr []T) (res []T, item T, ok bool) {
	if len(arr) < 1 {
		return
	}

	var t T
	le := len(arr) - 1
	res = arr[:le]
	item = arr[le]
	ok = true

	arr[le] = t

	return
}
```

</details>

#### PopFront

PopFront retrieves and deletes the element at the head of the slice


<details><summary>Code</summary>

```go

func PopFront[T any](arr []T) (res []T, item T, ok bool) {
	if len(arr) < 1 {
		res = arr
		return
	}

	item, res = arr[0], arr[1:]
	return
}
```

</details>

#### PushFront

PushFront inserts the item at the head of the slice


<details><summary>Code</summary>

```go

func PushFront[T any](arr []T, item T) []T {
	return append([]T{item}, arr...)
}
```

</details>

#### Shift

Shift inserts the item at the head of the slice


<details><summary>Code</summary>

```go

func Shift[T any](arr []T) ([]T, T, bool) {
	return PopFront(arr)
}
```

</details>

#### Unshift

Unshift inserts the item at the head of the slice


<details><summary>Code</summary>

```go

func Unshift[T any](arr []T, item T) []T {
	return PushFront(arr, item)
}
```

</details>



<br/>

### Maps

Package maps provides utilities to work with maps

Table of contents

- [Equals](####Equals)
- [Filter](####Filter)
- [FilterInPlace](####FilterInPlace)
- [FilterMap](####FilterMap)
- [FilterMapTuple](####FilterMapTuple)
- [Fold](####Fold)
- [Map](####Map)
- [Reduce](####Reduce)
- [Slice](####Slice)

#### Equals

Equals returns whether 2 maps are equals in values


<details><summary>Code</summary>

```go

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
```

</details>

#### Filter

Filter discards those entries from the map that do not match predicate.


<details><summary>Code</summary>

```go

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
```

</details>

#### FilterInPlace

FilterInPlace deletes those entries from the map that do not match predicate.


<details><summary>Code</summary>

```go

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
```

</details>

#### FilterMap

FilterMap both filters and maps a map. The predicate function should return a fp.Option monad:
fp.Some to indicate the entry should be kept.
fp.None to indicate the entry should be discarded


<details><summary>Code</summary>

```go

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
```

</details>

#### FilterMapTuple

FilterMapTuple both filters and maps the given map by receiving a predicate
which returns mapped values, and a boolean to indicate whether that entry
should be kept.


<details><summary>Code</summary>

```go

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
```

</details>

#### Fold

Fold compacts the given map into a single type by taking into account the initial value


<details><summary>Code</summary>

```go

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
```

</details>

#### Map

Map transforms a map into another one, with same or different types


<details><summary>Code</summary>

```go

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
```

</details>

#### Reduce

Reduce compacts the given map into a single type


<details><summary>Code</summary>

```go

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
```

</details>

#### Slice

Slice converts a map into a slice


<details><summary>Code</summary>

```go

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
```

</details>



<br/>

### Fp


Table of contents




<br/>

### Fp


Table of contents




<br/>

