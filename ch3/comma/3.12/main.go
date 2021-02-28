package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%t\n", isAnagram(os.Args[1], os.Args[2]))
}

//!+
// isAnagram returns whether args are anagrams or not .
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	char, size := utf8.DecodeRuneInString(a[0:])
	index := strings.Index(b, string(char))
	if index == -1 {
		return false
	}
	return isAnagram(a[size:], b[:index]+b[index+size:])
}

//!-
