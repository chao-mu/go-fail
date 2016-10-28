package fatal_test

import (
	"log"

	fatal "github.com/chao-mu/go-fatal"
)

func ExampleCheckOr() {
	// If run() returns an error, prints out and exits with "run faied: that error"
	// via a call to log.Fatal. One could also use logrus.Fatal and other
	// functions of the same signature.
	fatal.CheckOr(run(), log.Fatal, "run failed")
}

func ExampleCheck_JustError() {
	fatal.Check(run(), "")
}

func ExampleCheck() {
	// If run() returns an error, prints out "run failed: that error" and exits.
	fatal.Check(run(), "run failed")
}

func ExampleCheckf() {
	fatal.Checkf(run(), "Whoops '%v'. Lolz.", fatal.ErrorHere)
}

func run() error {
	return nil
}
