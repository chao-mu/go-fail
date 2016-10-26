package fatal_test

import (
	"errors"
	"os"
)

func ExampleCheckOr() {
	fatal.CheckOr(climbEverest(), log.Fatal, "Failed to climb everest")
}

func ExampleCheck() {
	// If climbEverest() returns an error, prints out "Aaaaaah: that error" and exits.
	fatal.Checkf(climbEverest(), "Aaaaaah")
}

func ExampleCheckf() {
	fatal.Checkf(climbEverest(), "Whoops '%v'. Lolz.", fatal.ErrorHere)
}
