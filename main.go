// Command canary is a tiny program the Fleet canary run works on.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: canary <command>")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "add":
		fmt.Println(Add(2, 3))
	default:
		fmt.Fprintf(os.Stderr, "canary: unknown command %q\n", os.Args[1])
		os.Exit(2)
	}
}

// Add returns the sum of two integers.
func Add(a, b int) int { return a + b }
