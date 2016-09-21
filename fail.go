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
	FailWithf(os.Stderr, 1, pattern, args...)
}

// FailWith writes the %v form of the provided value to the provided
// file handle ane exits with the provided status code.
func FailWith(file *os.File, code int, msg interface{}) {
	FailWithf(file, code, "%v", msg)
}

// FailWithf writes according to Printf rules to provided file handle
// and exits with the provided status code.
func FailWithf(file *os.File, code int, pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern, args...)
	os.Exit(code)
}
