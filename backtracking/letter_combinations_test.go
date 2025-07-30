package backtracking

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits   string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		actual := letterCombinations(test.digits)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("letterCombinations(%s) = %v; expected %v", test.digits, actual, test.expected)
		}
	}
}
