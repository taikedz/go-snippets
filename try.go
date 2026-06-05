package main

import (
	"fmt"
	"gocheat/snippets"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		readFiles(os.Args[1:])
	} else {
		readStdin()
	}
}

func Fail(status int, message string, items ...any) {
	fmt.Printf(message, items...)
	os.Exit(status)
}

func readFiles(filepaths []string) {
	for _, target := range filepaths {
		lines, err := snippets.ReadLines_file(target)
		if err != nil {
			Fail(1, "Error reading %v\n", err)
		}
		dump(lines)
	}
}

func readStdin() int {
	lines, err := snippets.ReadLines_stdin()
	if err != nil {
		Fail(1, "Stdin read fail: %v\n", err)
	}
	dump(lines)
	return 0
}

func dump(lines []string) {
	for _, line := range lines {
		fmt.Printf(": %s\n", line)
	}
}
