// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountByShifting returns the population count (number of set bits) of x.
func PopCountByShifting(x uint64) int {
	count := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) == 1 {
			count++
		}
	}
	return count
}

//!-
