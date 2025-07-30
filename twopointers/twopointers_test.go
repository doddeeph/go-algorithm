package twopointers

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}
	for _, test := range tests {
		actual := isPalindrome(test.s)
		if test.expected != actual {
			t.Errorf("isPalindrome(%s) = %v; expected %v", test.s, actual, test.expected)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, test := range tests {
		actual := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("twoSum(%v, %d) = %v; expected %v", test.nums, test.target, actual, test.expected)
		}
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, test := range tests {
		actual := threeSum(test.nums)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("threeSum(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}
