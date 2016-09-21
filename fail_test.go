package fail_test

import (
	"errors"
	"os"

	fail "github.com/chao-mu/go-fail"
)

func ExampleFail() {
	fail.Fail("I put all my money into Enron stocks")
}

func ExampleFailf() {
	fail.Failf("Discord detected: %v", errors.New("Aaaah!"))
}

func ExampleWith() {
	fail.With(os.Stdout, 23, "Hey there Standard out")
}

func ExampleWithf() {
	fail.Withf(os.Stdout, 23, "Almost there: %s", "Cool")
}

func ExampleCheck() {
	err := errors.New("uh oh...")
	fail.Check(err, "ummm: %v")
}
