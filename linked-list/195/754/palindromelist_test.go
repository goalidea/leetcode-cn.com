package palindromelist

import "testing"

func TestPalindromeList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	r := isPalindrome(head)

	if r {
		t.Error("determine palindrome list error")
	}
}
