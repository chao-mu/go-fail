A go library for when you just want to print to stderr and exit.
Don't you deserve to write less code?  Conditional exiting also supported for optimal laziness!

GoDoc: https://godoc.org/github.com/chao-mu/go-fatal

``` go
// Print "Discord detected" to stderr and exit
log.Fatal("Discord detected")
```

``` go
// If a non-nil error returns, print "Fatal error: " followed by that error.
log.Check(run(), "Fatal error")
```


``` go
// A more realistic example
file, err := os.Open(path)
fatal.Checkf(err, "Failed to open '%s': %v", path, fatal.ErrorHere)
```


# Logging framework integration

## Logrus
``` go
file, err := os.Open(path)
fatal.CheckOr(err, logrus.WithField("path", path).Fatal, "Failed to open file")
```

## log
``` go
file, err := os.Open(path)
fatal.CheckOrf(err, log.Fatal, "Failed to open '%s': %v", path, fatal.ErrorHere)
```
