A go library for when you just want to print to stderr and exit. Don't you deserve to write less code?

``` go
if err := run(); err != nil {
  fail.Failf("Discord detected: %v", err)
}
```

Or be even more lazy:

``` go
fail.Check(run(), "Discord detected: %v")
```

Documentation: https://godoc.org/github.com/chao-mu/go-fail
