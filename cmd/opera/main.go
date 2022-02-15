package main

import (
	"fmt"
	"os"

	"github.com/Fantom-foundation/go-opera/cmd/opera/launcher"

	// Force-load the tracer engines to trigger registration
	_ "github.com/Fantom-foundation/go-opera/tracers/js"
	_ "github.com/Fantom-foundation/go-opera/tracers/native"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
