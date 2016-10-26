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

// ErrorHere is used as a placeholder in Checkf
var ErrorHere = errors.New("Place holder error")

// Func represents a function that is expected to print/write its arguments
// and exit.
type Func func(v ...interface{})

// Check checks if err is non-nil and if so exits
// with the error and the provided prefix, or - if an empty
// prefix is specified - just the error.
func (f Func) Check(err error, prefix string) {
	if err == nil {
		return
	}

	if prefix == "" {
		f(err.Error())
	} else {
		f(fmt.Sprintf("%s: %v", prefix, err))
	}
}

// Check opperates as Func.Check does, except that it behaves like
// fmt.Sprintf.	Use ErrorHere in the argument list to specify
// where the error should be inserted.
func (f Func) Checkf(err error, format string, args ...interface{}) {
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

// Fatal prints to std err and exits with ErrorCode.
func Fatal(v ...interface{}) {
	fmt.Fprint(os.Stderr, fmt.Sprintln(v...))
	os.Exit(ErrorCode)
}

// Fatalf exits, first printing fmt.Sprintf(pattern, args) to stderr.
func Fatalf(pattern string, args ...interface{}) {
	Fatal(fmt.Sprintf(pattern, args...))
}

// Check operates like Func.Check, using Fatal as the print/exit function.
func Check(err error, prefix string) {
	CheckOr(err, Fatal, prefix)
}

// Checkf operates like Func.Checkf, using Fatal as the print/exit function.
func Checkf(err error, format string, args ...interface{}) {
	CheckOrf(err, Fatal, format, args...)
}

// CheckOr is a convenience function that allows you to specify your own 
// fatal function. Handy for integrating with logging frameworks.
// It is equivalent to fatal.Func(f).Check.
func CheckOr(err error, f Func, prefix string) {
	f.Check(err, prefix)
}

// CheckOrf is a convenience function that is equivalent
// to fatal.Func(f).Checkf.
func CheckOrf(err error, f Func, prefix string, args ...interface{}) {
	f.Checkf(err, prefix, args...)
}
