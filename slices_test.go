package stadio

import (
	"testing"
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
