package recursion

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, test := range tests {
		actual := generateParentheses(test.n)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("generateParentheses(%d) = %v; expected %v", test.n, actual, test.expected)
		}
	}
}
