// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import "testing"

func Test_comma(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"1234567890", "1,234,567,890"},
	}
	for _, tt := range tests {
		if got := comma(tt.s); got != tt.want {
			t.Errorf("comma() = %v, want %v", got, tt.want)
		}
	}
}
