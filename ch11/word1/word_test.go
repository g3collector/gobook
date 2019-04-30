package word1

import (
	"testing"
)

func TestName(t *testing.T) {
	t.Log("hey")
}

func TestSin(t *testing.T) {

}

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartraded") {
		t.Error(`IsPalindrome("detartraded") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
