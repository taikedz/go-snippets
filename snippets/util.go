package snippets

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func ArrayHas(term string, stuff []string) bool {
	for _, thing := range stuff {
		if term == thing {
			return true
		}
	}
	return false
}

func ExcludeStr(input []string, exclude []string) []string {
	var retained []string

	for _, s := range input {
		if !ArrayHas(s, exclude) {
			retained = append(retained, s)
		}
	}

	return retained
}

func ExtractValueOfKey(key string, items []string) (string, error) {
	// assume an array of "key=value" strings
	// locate key , split on '=', return the value
	key_eq := fmt.Sprintf("%s=", key)
	for _, item := range items {
		if strings.Index(item, key_eq) == 0 {
			return item[len(key_eq):], nil
		}
	}
	return "", fmt.Errorf("required parameter '%s' not found", key)
}

func IsRootUser() bool {
	if os.Getenv("PAF_TEST_PMAN") != "" {
		// test mode - we'll always say we're not root, to catch sudo detection
		return false
	}
	u, e := user.Current()
	FailIf(e, 98, "Fatal - Could not get current user!")
	return u.Uid == "0" // posix only!
}

func IsWinAdmin() (bool, error) {
	/* This is apparently the way to handle Windows.
	   For now, not supporting windows choco/winget
	   But this could be a target for future
	*/

	// https://stackoverflow.com/a/19847868/2703818
	if runtime.GOOS != "windows" {
		return false, fmt.Errorf("not on Windows")
	}

	// https://stackoverflow.com/a/59147866/2703818
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false, nil
	}
	return true, nil
}

func SplitStringMultichar(data string, chars string) []string {
	tokens := []string{data}

	for _, c := range chars {
		tokens = SplitStringsChar(tokens, string(c))
	}
	return tokens
}

func SplitStringsChar(data []string, char string) []string {
	var tokens []string
	for _, piece := range data {
		tokens = append(tokens, strings.Split(piece, string(char))...)
	}
	return tokens
}
