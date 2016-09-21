package fail

import (
	"fmt"
	"os"
)

// Fail prints the %v form of the provided value to stderr and exits with 1.
func Fail(msg interface{}) {
	Failf("%v", msg)
}

// Failf prints according to Printf rules to stderr and exits with 1.
func Failf(pattern string, args ...interface{}) {
	FailWithf(1, pattern, args...)
}

// FailWith prints the %v form of the provided value to stderr and exits
// with provided status code.
func FailWith(code int, msg interface{}) {
	FailWithf(code, "%v", msg)
}

// FailWithf prints according to Printf rules to stderr and exits with
// provided status code.
func FailWithf(code int, pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern, args...)
	os.Exit(code)
}
