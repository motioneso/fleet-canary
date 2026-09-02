// Command canary is a tiny program the Fleet canary run works on.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: canary <command>")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "add":
		fmt.Println(Add(2, 3))
	case "greet":
		name := ""
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		fmt.Println(Greet(name))
	case "farewell":
		name := ""
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		fmt.Println(Farewell(name))
	default:
		fmt.Fprintf(os.Stderr, "canary: unknown command %q\n", os.Args[1])
		os.Exit(2)
	}
}

// Add returns the sum of two integers.
func Add(a, b int) int { return a + b }

// Greet returns the greeting line for name, without a trailing newline.
func Greet(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "Hello, world!"
	}
	return "Hello, " + trimmed + "!"
}

// Farewell returns the farewell line for name, without a trailing newline.
func Farewell(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "Goodbye, world!"
	}
	return "Goodbye, " + trimmed + "!"
}
