package ggassert_test

import (
	"fmt"
	"testing"

	"github.com/thara/ggassert"
)

func TestEqual(t *testing.T) {
	type pattern[T any] struct {
		expected, actual T
		failed           bool
	}

	t.Run("int", func(t *testing.T) {
		tests := []pattern[int]{
			{2, 3, true},
			{9, 9, false},
			{10, 1, true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.Equal(mock, tt.expected, tt.actual, "failed")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		type sample struct {
			a, b, c int
			s       string
		}
		tests := []pattern[sample]{
			{sample{}, sample{}, false},
			{sample{}, sample{1, 2, 3, "test"}, true},
			{sample{1, 2, 3, "test"}, sample{1, 2, 3, "test"}, false},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.Equal(mock, tt.expected, tt.actual, "failed")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

type orderedPattern[T any] struct {
	op1, op2 T
	failed   bool
}

func TestLessThan(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []orderedPattern[int]{
			{8, 9, false},
			{9, 9, true},
			{10, 9, true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.LessThan(mock, tt.op1, tt.op2, "failed")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func TestLessThanOrEqual(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []orderedPattern[int]{
			{8, 9, false},
			{9, 9, false},
			{10, 9, true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.LessThanOrEqual(mock, tt.op1, tt.op2, "failed")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []orderedPattern[int]{
			{8, 9, true},
			{9, 9, true},
			{10, 9, false},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.GreaterThan(mock, tt.op1, tt.op2, "failed")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func TestGreaterThanOrEqual(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []orderedPattern[int]{
			{8, 9, true},
			{9, 9, false},
			{10, 9, false},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.GreaterThanOrEqual(mock, tt.op1, tt.op2, "failed")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func TestContainsSlice(t *testing.T) {
	type pattern[T comparable] struct {
		slice   []T
		element T
		failed  bool
	}

	type sample struct {
		a, b, c int
		s       string
	}

	t.Run("struct", func(t *testing.T) {
		tests := []pattern[sample]{
			{[]sample{}, sample{1, 2, 3, "test"}, true},
			{[]sample{
				{1, 2, 3, "test"},
				{2, 2, 3, "test"},
			}, sample{1, 2, 3, "test"}, false},
			{[]sample{
				{2, 2, 3, "test"},
				{1, 2, 3, "test2"},
			}, sample{1, 2, 3, "test"}, true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.ContainsSlice(mock, tt.slice, tt.element, "fail")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})

	t.Run("struct pointer", func(t *testing.T) {
		tests := []pattern[*sample]{
			{[]*sample{}, &sample{1, 2, 3, "test"}, true},
			{[]*sample{
				{1, 2, 3, "test"},
				{2, 2, 3, "test"},
			}, &sample{1, 2, 3, "test"}, false},
			{[]*sample{
				{2, 2, 3, "test"},
				{1, 2, 3, "test2"},
			}, &sample{1, 2, 3, "test"}, true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.ContainsSlice(mock, tt.slice, tt.element, "fail")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func TestContainsMapKey(t *testing.T) {
	type pattern[K comparable, V any] struct {
		m      map[K]V
		key    K
		failed bool
	}

	t.Run("string", func(t *testing.T) {
		tests := []pattern[string, int]{
			{map[string]int{}, "aaa", true},
			{map[string]int{
				"aaa": 1,
				"bbb": 2,
				"ccc": 3,
			}, "aaa", false},
			{map[string]int{
				"bbb": 2,
				"ccc": 3,
				"ddd": 4,
			}, "aaa", true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.ContainsMapKey(mock, tt.m, tt.key, "fail")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func TestContainsMapValue(t *testing.T) {
	type pattern[K comparable, V any] struct {
		m      map[K]V
		value  V
		failed bool
	}

	t.Run("string", func(t *testing.T) {
		tests := []pattern[int, string]{
			{map[int]string{}, "aaa", true},
			{map[int]string{
				111: "aaa",
				222: "bbb",
				333: "ccc",
			}, "aaa", false},
			{map[int]string{
				222: "bbb",
				333: "ccc",
				444: "ddd",
			}, "aaa", true},
		}
		for i, tt := range tests {
			t.Run(fmt.Sprintf("pattern:%d", i), func(t *testing.T) {
				mock := new(testing.T)
				ggassert.ContainsMapValue(mock, tt.m, tt.value, "fail")
				if mock.Failed() != tt.failed {
					t.Errorf("unexpected assertion result. expected %t, but actual %t", tt.failed, mock.Failed())
				}
			})
		}
	})
}

func Example() {
	t := new(testing.T)

	ggassert.Equal(t, 2, 2, "failed") // pass
	ggassert.Equal(t, 1, 2, "failed") // failed
	// ggassert.Equal(t, 1, "aaa", "failed") // compile error: default type string of "aaa" does not match inferred type int for T

	ggassert.LessThan(t, 1, 2, "failed")        // pass
	ggassert.LessThan(t, 2, 2, "failed")        // failed
	ggassert.LessThanOrEqual(t, 2, 2, "failed") // pass
	// ggassert.LessThanOrEqual(t, 2, 2.0, "failed") // compile error: default type float64 of 2.0 does not match inferred type int for T

	ggassert.GreaterThan(t, 2, 1, "failed")        // pass
	ggassert.GreaterThan(t, 2, 2, "failed")        // failed
	ggassert.GreaterThanOrEqual(t, 2, 2, "failed") // pass
	// ggassert.GreaterThanOrEqual(t, 2, 2.0, "failed") // compile error: default type float64 of 2.0 does not match inferred type int for T

	ggassert.ContainsSlice(t, []int{1, 2, 3}, 2, "failed") // pass
	ggassert.ContainsSlice(t, []int{1, 2, 3}, 4, "failed") // failed
	// ggassert.ContainsSlice(t, []int{1, 2, 3}, "aaa", "failed") // compile error: cannot use "aaa" (untyped string constant) as int value in argument to ggassert.ContainsSlice

	m := map[string]int{
		"aaa": 111,
		"bbb": 222,
		"ccc": 333,
	}

	ggassert.ContainsMapKey(t, m, "aaa", "failed") // pass
	ggassert.ContainsMapKey(t, m, "ddd", "failed") // failed
	ggassert.ContainsMapValue(t, m, 111, "failed") // pass
	ggassert.ContainsMapValue(t, m, 444, "failed") // failed

	// ggassert.ContainsMapKey(t, m, 111, "failed")     // compile error: cannot use 111 (untyped int constant) as string value in argument to ggassert.ContainsMapKey
	// ggassert.ContainsMapValue(t, m, "aaa", "failed") // compile error: cannot use "aaa" (untyped string constant) as int value in argument to ggassert.ContainsMapValue

	// Output:
}
