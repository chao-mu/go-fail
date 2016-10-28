// Package fatal simplifies the process of printing an error message
// and then exiting. Great for prototyping or long term development
// with integration into logging frameworks. (For example, you can
// pass log.Fatal to CheckOr).
package fatal

import (
	"errors"
	"fmt"
	"os"
)

const (
	// ErrorCode is the status code we return on exit.
	ErrorCode = 1
)

// ErrorHere is used as a placeholder in Checkf/CheckOrf  .
var ErrorHere = errors.New("Place holder error")

// Fatal prints to std err and exits with ErrorCode.
func Fatal(v ...interface{}) {
	fmt.Fprint(os.Stderr, fmt.Sprintln(v...))
	os.Exit(ErrorCode)
}

// Fatalf exits, first printing fmt.Sprintf(pattern, args) to stderr.
func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}

// Check checks if err is non-nil and if so exits
// with the error and the provided prefix, or - if an empty
// prefix is specified - just the error.
func Check(err error, prefix string) {
	CheckOr(err, Fatal, prefix)
}

// Checkf opperates as Check does, except that it behaves like
// fmt.Sprintf.	Use ErrorHere in the argument list to specify
// where the error should be inserted.
func Checkf(err error, format string, args ...interface{}) {
	CheckOrf(err, Fatal, format, args...)
}

// CheckOr is a replacement for Check when you want to provide your
// own fatal function. Handy for integrating with logging frameworks.
func CheckOr(err error, f func(...interface{}), prefix string) {
	if prefix == "" {
		CheckOrf(err, f, "%s", err)
	} else {
		CheckOrf(err, f, "%s: %v", prefix, err)
	}
}

// CheckOrf is a replacement for Checkf and functions the same way,
// except using the provided function to print and exit.
func CheckOrf(err error, f func(...interface{}), format string, args ...interface{}) {
	if err == nil {
		return
	}

	for i, arg := range args {
		if arg == ErrorHere {
			args[i] = err
		}
	}

	f(fmt.Sprintf(format, args...))
}
