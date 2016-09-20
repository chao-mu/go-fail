A go library for when you just want to print to stderr and exit. Don't you deserve one less line of code?

``` go
if err := run(); err != nil {
  fail.Failf("Discord detected: %v", err)
}
```

Documentation: https://godoc.org/github.com/chao-mu/go-fail
