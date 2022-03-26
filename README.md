# gg-assert

`gg-assert` is a deadly simple & easy assertion library for golang.

It's also using generics supported since Go 1.18 later.

## Usage

```
ggassert.Equal(t, 1, 2, "failed")
ggassert.Equal(t, 1, "aaa", "failed") // compile error: default type string of "aaa" does not match inferred type int for T
```

## Installation

```
go get github.com/thara/ggassert
```

## License

MIT

## Author

Tomochika Hara (a.k.a [thara](https://thara.dev))
