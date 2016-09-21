// Package fail simplifies the process of printing/writing an
// error message and then exiting.
package fail

import (
	"fmt"
	"io"
	"os"
)

const (
	// Pattern dictates how Fail/With print the provided value.
	Pattern = "%v"
	// ErrorCode is the status code we return on exit
	ErrorCode = 1
)

// format applys Pattern to msg via sprintf.
func format(msg interface{}) string {
	return fmt.Sprintf(Pattern, msg)
}

// Fail writes msg to stderr (according to fail.Pattern),
// then exits with fail.ErrorCode.
func Fail(msg interface{}) {
	Failf(Pattern, msg)
}

// Failf writes pattern to stderr parameterized by args according to printf
// then exits with 1. A newline is appended.
func Failf(pattern string, args ...interface{}) {
	Withf(os.Stderr, ErrorCode, pattern, args...)
}

// With writes msg to provided writer (according to fail.Pattern)
// and exits with provided code.
func With(w io.Writer, code int, msg interface{}) {
	Withf(w, code, Pattern, msg)
}

// Withf writes pattern to the provided writer parameterized by args
// according to printf then exits with provided code. A newline is appended.
func Withf(w io.Writer, code int, pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern+"\n", args...)
	os.Exit(code)
}

// Check checks if the error passed is non-nil, and if it is
// delegates to Failf with the provided printf format and err
// as its first arg.
func Check(err error, pattern string) {
	if err != nil {
		Failf(pattern, err)
	}
}

// CheckWith checks if the error passed is non-nil, and
// if it is delegates to Withf with the provided printf format
// and err as its first arg.
func CheckWith(err error, w io.Writer, code int, pattern string) {
	if err != nil {
		Withf(w, code, pattern, err)
	}
}
