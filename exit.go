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
	ExitWithf(1, pattern, args...)
}

// ExitWith prints the %v form of the provided value to stderr and exits
// with provided status code.
func ExitWith(code int, msg interface{}) {
	ExitWithf(code, "%v", msg)
}

// ExitWithf prints according to Printf rules to stderr and exits with
// provided status code.
func ExitWithf(code int, pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern, args...)
	os.Exit(code)
}
