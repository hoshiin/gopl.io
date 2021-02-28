package main

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"abc", "cba", true},
		{"aabb", "abab", true},
		{"aAbBcC", "abcABC", true},
		{"aAbBcC", "abcABCD", false},
		{"デンセツノケン", "ケセノンツンデ", true},
		{"デンセツノケン", "ケセンツンデ", false},
		{"デンセツノケン", "ケセンツンデデ", false},
	}
	for _, tt := range tests {
		if got := isAnagram(tt.a, tt.b); got != tt.want {
			t.Errorf("isAnagram() = %v, want %v", got, tt.want)
		}
	}
}
