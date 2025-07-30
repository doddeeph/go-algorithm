package stack

import (
	"reflect"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
	}
	for _, test := range tests {
		actual := isValidParentheses(test.s)
		if test.expected != actual {
			t.Errorf("isValidParentheses(%s) = %v; expected %v", test.s, actual, test.expected)
		}
	}
}

func TestMinStack(t *testing.T) {
	tests := []struct {
		cmds     []string
		vals     [][]int
		expected []int
	}{
		{[]string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			[][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			[]int{0, 0, 0, 0, -3, 0, 0, -2}},
	}
	for _, test := range tests {
		actual := minStack(test.cmds, test.vals)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("minStack(%v, %v) = %v; expeceted %v", test.cmds, test.vals, actual, test.expected)
		}
	}
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected int
	}{
		// ((2 + 1) * 3) = 9
		{[]string{"2", "1", "+", "3", "*"}, 9},
		// (4 + (13 / 5)) = 6
		{[]string{"4", "13", "5", "/", "+"}, 6},
		// // ((10 * (6 / ((9 + 3) * -11))) + 17) + 5 = 22
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, test := range tests {
		actual := evalRPN(test.tokens)
		if test.expected != actual {
			t.Errorf("evalRPN(%v) = %d; expected %d", test.tokens, actual, test.expected)
		}
	}
}
