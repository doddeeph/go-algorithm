package slidingwindow

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		actual := maxProfit(test.prices)
		if test.expected != actual {
			t.Errorf("maxProfit(%v) = %d; expected %d", test.prices, actual, test.expected)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, test := range tests {
		actual := lengthOfLongestSubstring(test.s)
		if test.expected != actual {
			t.Errorf("lengthOfLongestSubstring(%s) = %d; expected %d", test.s, actual, test.expected)
		}
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct{
		s string
		k int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}
	for _, test := range tests {
		actual := characterReplacement(test.s, test.k)
		if test.expected != actual {
			t.Errorf("characterReplacement(%s, %d) = %d; expected %d", test.s, test.k, actual, test.expected)
		}
	}
}