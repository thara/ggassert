package ggassert

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

// Equal asserts that two any values are equal.
func Equal[T any](t testing.TB, expected, actual T, format string, args ...any) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	t.Errorf(format, args...)
}

// LessThan asserts that the first operand is less than the second one
func LessThan[T constraints.Ordered](t testing.TB, a, b T, format string, args ...any) {
	if a < b {
		return
	}
	t.Errorf(format, args...)
}

// LessThanOrEqual asserts that the first operand is less than or equal the second one
func LessThanOrEqual[T constraints.Ordered](t testing.TB, a, b T, format string, args ...any) {
	if a <= b {
		return
	}
	t.Errorf(format, args...)
}

// GreaterThan asserts that the first operand is greater than the second one
func GreaterThan[T constraints.Ordered](t testing.TB, a, b T, format string, args ...any) {
	if a > b {
		return
	}
	t.Errorf(format, args...)
}

// GreaterThanOrEqual asserts that the first operand is greater than or equal the second one
func GreaterThanOrEqual[T constraints.Ordered](t testing.TB, a, b T, format string, args ...any) {
	if a >= b {
		return
	}
	t.Errorf(format, args...)
}

// ContainsSlice asserts that the slice contains the specified value.
func ContainsSlice[T comparable](t testing.TB, s []T, expected T, format string, args ...any) {
	for _, v := range s {
		if reflect.DeepEqual(expected, v) {
			return
		}
	}
	t.Errorf(format, args...)
}

// ContainsMapKey asserts that the map contains the specified key.
func ContainsMapKey[K comparable, V any](t testing.TB, target map[K]V, expectedKey K, format string, args ...any) {
	_, ok := target[expectedKey]
	if ok {
		return
	}
	t.Errorf(format, args...)
}

// ContainsMapValue asserts that the map contains the specified value.
func ContainsMapValue[K, V comparable](t testing.TB, target map[K]V, expectedValue V, format string, args ...any) {
	for _, v := range target {
		if reflect.DeepEqual(v, expectedValue) {
			return
		}
	}
	t.Errorf(format, args...)
}
