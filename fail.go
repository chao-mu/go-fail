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

/*
type Fail struct {
	Code int
	Out  io.Writer
}

// Exitf writes msg to stderr (according to fail.Pattern),
// then exits with .
func (f Fail) Exit(msg interface{}) {
	f.Exitf(Pattern, msg)
}
*/

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
// delegates to Failf, with the error and provided prefix
// in the form "%s: %v" (or falls through to Fail if no prefix).
func Check(err error, prefix string) {
	if err == nil {
		return
	}

	if prefix == "" {
		Fail(err)
	} else {
		Failf("%s: %v", prefix, err)
	}
}

// CheckWith is like Check, except that the writer and exit
// code can be specified.
func CheckWith(err error, prefix string, w io.Writer, code int) {
	if err == nil {
		return
	}

	if prefix == "" {
		Withf(w, code, "%s: %v", prefix, err)
	} else {
		With(w, code, err)
	}
}
