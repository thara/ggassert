# gg-assert

`gg-assert` is a deadly simple & easy assertion library for golang.

It's also using generics supported since Go 1.18 later.

## Usage

```go
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
```

## Installation

```
go get github.com/thara/ggassert
```

## License

MIT

## Author

Tomochika Hara (a.k.a [thara](https://thara.dev))
