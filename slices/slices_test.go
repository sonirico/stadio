package slices

import (
	"testing"

	"github.com/sonirico/stadio/fp"
)

func TestSlice_Len(t *testing.T) {
	type testCase struct {
		name           string
		payload        Slice[int]
		expectedLength int
	}

	tests := []testCase{
		{
			name:           "zero length slice",
			payload:        Slice[int]([]int{}),
			expectedLength: 0,
		},
		{
			name:           "nil slice",
			payload:        Slice[int](nil),
			expectedLength: 0,
		},
		{
			name:           "slice with more than one",
			payload:        Slice[int]([]int{1, 2, 3}),
			expectedLength: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expectedLength != test.payload.Len() {
				t.Errorf("unexpected slice length. want %d, have %d",
					test.expectedLength, test.payload.Len())
			}
		})
	}
}

func TestSlice_Range(t *testing.T) {
	type testCase struct {
		name           string
		payload        Slice[int]
		expectedLength int
	}

	tests := []testCase{
		{
			name:           "zero length slice",
			payload:        Slice[int]([]int{}),
			expectedLength: 0,
		},
		{
			name:           "nil slice",
			payload:        Slice[int](nil),
			expectedLength: 0,
		},
		{
			name:           "slice with more than one",
			payload:        Slice[int]([]int{1, 2, 3}),
			expectedLength: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualLen := 0
			test.payload.Range(func(x int, _ int) bool {
				actualLen += 1
				return true
			})
			if test.expectedLength != actualLen {
				t.Errorf("unexpected slice length. want %d, have %d",
					test.expectedLength, actualLen)
			}
		})
	}
}

func TestSlice_Range_EarlyReturn(t *testing.T) {
	slice := Slice[int]([]int{1, 2, 3})
	actualLen := 0
	expectedLength := 2
	slice.Range(func(x int, i int) bool {
		actualLen += 1
		return i%2 == 0
	})
	if actualLen != expectedLength {
		t.Errorf("unexpected length. want %d, have %d", expectedLength, actualLen)
	}
}

func TestSlice_Get(t *testing.T) {
	type testCase struct {
		name        string
		payload     Slice[int]
		index       int
		expectedOk  bool
		expectedRes int
	}

	tests := []testCase{
		{
			name:        "negative index",
			payload:     Slice[int]([]int{}),
			index:       -1,
			expectedRes: 0,
			expectedOk:  false,
		},
		{
			name:        "zero index",
			payload:     Slice[int]([]int{1, 2, 3}),
			index:       0,
			expectedRes: 1,
			expectedOk:  true,
		},
		{
			name:        "zero index for nil slice",
			payload:     Slice[int](nil),
			index:       0,
			expectedRes: 0,
			expectedOk:  false,
		},
		{
			name:        "index of last item",
			payload:     Slice[int]([]int{1, 2, 3}),
			index:       2,
			expectedRes: 3,
			expectedOk:  true,
		},
		{
			name:        "out of bounds index",
			payload:     Slice[int]([]int{1, 2, 3}),
			index:       3,
			expectedRes: 0,
			expectedOk:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actualRes, actualOk := test.payload.Get(test.index)

			if test.expectedOk != actualOk {
				t.Errorf("unexpected ok, want %t, have %t", test.expectedOk, actualOk)
			}
			if test.expectedRes != actualRes {
				t.Errorf("unexpected value, want %d, have %d", test.expectedRes, actualRes)
			}
		})
	}
}

func TestSlice_IndexOf(t *testing.T) {
	type testCase struct {
		name        string
		payload     Slice[int]
		predicate   func(i int) bool
		expectedIdx int
	}

	tests := []testCase{
		{
			name:    "nil slice should return -1",
			payload: Slice[int]([]int{}),
			predicate: func(i int) bool {
				return true
			},
			expectedIdx: -1,
		},
		{
			name:    "item at the first position",
			payload: Slice[int]([]int{1, 2, 3}),
			predicate: func(i int) bool {
				return i == 1
			},
			expectedIdx: 0,
		},
		{
			name:    "item at the last position",
			payload: Slice[int]([]int{1, 2, 3}),
			predicate: func(i int) bool {
				return 3 == i
			},
			expectedIdx: 2,
		},
		{
			name:    "item not found",
			payload: Slice[int]([]int{73, 30, 5}),
			predicate: func(i int) bool {
				return 42 == i
			},
			expectedIdx: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualIdx := test.payload.IndexOf(test.predicate)

			if test.expectedIdx != actualIdx {
				t.Errorf("unexpected value, want %d, have %d", test.expectedIdx, actualIdx)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type testCase struct {
		name     string
		payload  Slice[int]
		target   int
		expected bool
	}

	tests := []testCase{
		{
			name:     "nil slice should return false",
			payload:  Slice[int]([]int{}),
			target:   1,
			expected: false,
		},
		{
			name:     "item at the first position",
			payload:  Slice[int]([]int{1, 2, 3}),
			target:   1,
			expected: true,
		},
		{
			name:     "item at the last position",
			payload:  Slice[int]([]int{1, 2, 3}),
			target:   3,
			expected: true,
		},
		{
			name:     "item not found",
			payload:  Slice[int]([]int{73, 30, 5}),
			target:   3,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualContains := Contains(test.payload, func(x int) bool { return x == test.target })

			if test.expected != actualContains {
				t.Errorf("unexpected value, want %t, have %t", test.expected, actualContains)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type testCase struct {
		name      string
		payload   Slice[int]
		expected  Slice[int]
		predicate func(int) bool
	}

	tests := []testCase{
		{
			name:     "nil slice should return nil slice",
			payload:  Slice[int]([]int{}),
			expected: Slice[int]([]int{}),
			predicate: func(i int) bool {
				return true
			},
		},
		{
			name:     "elements are filtered leaving some",
			payload:  Slice[int]([]int{1, 2, 3}),
			expected: Slice[int]([]int{2}),
			predicate: func(i int) bool {
				return i%2 == 0
			},
		},
		{
			name:     "elements are filtered leaving none",
			payload:  Slice[int]([]int{1, 2, 3}),
			expected: Slice[int]([]int{}),
			predicate: func(i int) bool {
				return i > 10
			},
		},
	}

	for _, test := range tests {
		t.Run("[Filter] "+test.name, func(t *testing.T) {
			actual := Filter(test.payload.Clone(), test.predicate)

			if !test.expected.Equals(actual, func(x, y int) bool {
				return x == y
			}) {
				t.Errorf("unexpected value, want %v, have %v", test.expected, actual)
			}
		})

		t.Run("[FilterInPlace] "+test.name, func(t *testing.T) {
			actual := FilterInPlace(test.payload.Clone(), test.predicate)

			if !test.expected.Equals(actual, func(x, y int) bool {
				return x == y
			}) {
				t.Errorf("unexpected value, want %v, have %v", test.expected, actual)
			}
		})

		t.Run("[FilterInPlaceCopy] "+test.name, func(t *testing.T) {
			actual := FilterInPlaceCopy(test.payload.Clone(), test.predicate)

			if !test.expected.Equals(actual, func(x, y int) bool {
				return x == y
			}) {
				t.Errorf("unexpected value, want %v, have %v", test.expected, actual)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	type testCase struct {
		name      string
		payload   Slice[int]
		expected  Slice[int]
		predicate func(int) fp.Option[int]
	}

	tests := []testCase{
		{
			name:     "nil slice should return nil slice",
			payload:  Slice[int]([]int{}),
			expected: Slice[int]([]int{}),
			predicate: func(i int) fp.Option[int] {
				return fp.None[int]()
			},
		},
		{
			name:     "elements are filtered leaving some",
			payload:  Slice[int]([]int{1, 2, 3}),
			expected: Slice[int]([]int{4}),
			predicate: func(i int) fp.Option[int] {
				if i%2 == 0 {
					return fp.Some(i * i)
				}
				return fp.None[int]()
			},
		},
		{
			name:     "elements are filtered leaving none",
			payload:  Slice[int]([]int{1, 2, 3}),
			expected: Slice[int]([]int{}),
			predicate: func(i int) fp.Option[int] {
				return fp.None[int]()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := FilterMap(test.payload, test.predicate)

			if !test.expected.Equals(actual, func(x, y int) bool {
				return x == y
			}) {
				t.Errorf("unexpected value, want %v, have %v", test.expected, actual)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type testCase struct {
		name     string
		payload  Slice[int]
		expected int
	}

	predicate := func(x, y int) int { return x + y }

	tests := []testCase{
		{
			name:     "nil slice should return zero value",
			payload:  Slice[int]([]int{}),
			expected: 0,
		},
		{
			name:     "slice with only 1 element",
			payload:  Slice[int]([]int{1}),
			expected: 1,
		},
		{
			name:     "slice with several elements",
			payload:  Slice[int]([]int{1, 2, 3}),
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ReduceSame(test.payload, predicate)

			if test.expected != actual {
				t.Errorf("unexpected value, want %d, have %d", test.expected, actual)
			}
		})
	}
}

func TestCut(t *testing.T) {
	type testCase struct {
		name     string
		payload  Slice[int]
		from     int
		to       int
		expected Slice[int]
	}

	tests := []testCase{
		{
			name:     "nil slice should be noop",
			payload:  Slice[int]([]int{}),
			from:     0,
			to:       0,
			expected: Slice[int]([]int{}),
		},
		{
			name:     "slice with one item",
			payload:  Slice[int]([]int{1}),
			from:     0,
			to:       0,
			expected: Slice[int]([]int{}),
		},
		{
			name:     "slice with two items cut first one",
			payload:  Slice[int]([]int{1, 2}),
			from:     0,
			to:       0,
			expected: Slice[int]([]int{2}),
		},
		{
			name:     "slice with two items cut last one",
			payload:  Slice[int]([]int{1, 2}),
			from:     1,
			to:       1,
			expected: Slice[int]([]int{1}),
		},
		{
			name:     "slice with two items cut all",
			payload:  Slice[int]([]int{1, 2}),
			from:     0,
			to:       1,
			expected: Slice[int]([]int{}),
		},
		{
			name:     "`from` greater than to should consider `to` to be amount",
			payload:  Slice[int]([]int{1, 2}),
			from:     1,
			to:       0,
			expected: Slice[int]([]int{1}),
		},
		{
			name:     "`from` greater than slice length is moved to end",
			payload:  Slice[int]([]int{1, 2}),
			from:     3,
			to:       0,
			expected: Slice[int]([]int{1}),
		},
		{
			name:     "`to` greater than slice length is moved to end",
			payload:  Slice[int]([]int{1, 2}),
			from:     0,
			to:       3,
			expected: Slice[int]([]int{}),
		},
		{
			name:     "`from` lower than zero is moved to zero",
			payload:  Slice[int]([]int{1, 2}),
			from:     -1,
			to:       0,
			expected: Slice[int]([]int{2}),
		},
		{
			name:     "`to` lower than zero is moved to zero",
			payload:  Slice[int]([]int{1, 2}),
			from:     0,
			to:       -1,
			expected: Slice[int]([]int{2}),
		},
		{
			name:     "cut with more than two items cut all",
			payload:  Slice[int]([]int{1, 2, 3, 4}),
			from:     0,
			to:       4,
			expected: Slice[int]([]int{}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Cut(test.payload, test.from, test.to)

			if !test.expected.Equals(actual, func(x, y int) bool { return x == y }) {
				t.Errorf("unexpected value, want %v, have %v", test.expected, actual)
			}
		})
	}
}
