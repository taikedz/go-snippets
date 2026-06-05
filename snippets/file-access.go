package snippets

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines_stdin() ([]string, error) {
	fileScanner := bufio.NewScanner(os.Stdin)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		if err := fileScanner.Err(); err != nil {
			return nil, err
		}
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func ReadLines_file(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}
