package word2

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartraded", true},
		{"A man, a plan, a canal: Panama", true},
	}

	for _, test := range tests {
		fmt.Println("hey")
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("Is Palindrome(%q) = %v", test.input, got)
		}
	}
}
