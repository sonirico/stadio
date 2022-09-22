package maps

import (
	"strconv"
	"testing"

	"github.com/sonirico/stadio/fp"
	"github.com/sonirico/stadio/tuples"
)

func TestMap(t *testing.T) {
	type (
		testCase struct {
			name     string
			payload  map[int]int
			expected map[string]string
		}
	)

	tests := []testCase{
		{
			name:     "nil map is noop",
			payload:  nil,
			expected: nil,
		},
		{
			name:     "empty map returns empty map",
			payload:  map[int]int{},
			expected: map[string]string{},
		},
		{
			name:     "filled map",
			payload:  map[int]int{100: 3, 29: 2},
			expected: map[string]string{"100": "9", "29": "4"},
		},
	}

	predicate := func(k, v int) (string, string) {
		return strconv.FormatInt(int64(k), 10), strconv.FormatInt(int64(v*v), 10)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Map(test.payload, predicate)

			if !Equals(test.expected, actual, assertMapValueEq) {
				t.Errorf("unexpected map\nwant %v\nhave %v",
					test.expected, actual)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	type (
		testCase struct {
			name     string
			payload  map[int]int
			expected map[string]string
		}
	)

	tests := []testCase{
		{
			name:     "nil map is noop",
			payload:  nil,
			expected: nil,
		},
		{
			name:     "empty map returns empty map",
			payload:  map[int]int{},
			expected: map[string]string{},
		},
		{
			name:     "filled map",
			payload:  map[int]int{101: 3, 22: 2},
			expected: map[string]string{"101": "9"},
		},
	}

	predicate := func(k, v int) fp.Option[tuples.Tuple2[string, string]] {
		if k%2 == 0 {
			return fp.None[tuples.Tuple2[string, string]]()
		}

		return fp.Some(tuples.Tuple2[string, string]{
			strconv.FormatInt(int64(k), 10),
			strconv.FormatInt(int64(v*v), 10),
		})
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := FilterMap(test.payload, predicate)

			if !Equals(test.expected, actual, assertMapValueEq) {
				t.Errorf("unexpected map\nwant %v\nhave %v",
					test.expected, actual)
			}
		})
	}
}

func assertMapValueEq(x, y string) bool {
	return x == y
}
