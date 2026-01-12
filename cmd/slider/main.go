package main

import (
	"fmt"
	"os"

	"github.com/vysmv/slider/internal/app"
	"github.com/vysmv/slider/internal/store"
	"github.com/vysmv/slider/internal/ui/term"
)

func main() {
	if len(os.Args) < 2 {
		// Write an error message to the standard error stream (stderr)
		fmt.Fprintln(os.Stderr, "usage: slider <slides_dir>")
		os.Exit(1)
	}

	dir := os.Args[1]

	st, err := store.NewDirStore(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	ui := term.NewUI()
	a := app.New(st, ui)

	if err := a.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
}
