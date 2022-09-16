
# Stadio

Compendium of functions, data structures, monadic wrappers and more which, hopefully, will be
included as a standard library of the language

## Modules

### slices
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



<br/>### fp

Table of contents




<br/>