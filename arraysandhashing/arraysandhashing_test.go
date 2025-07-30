package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, test := range tests {
		actual := containsDuplicate(test.nums)
		if actual != test.expected {
			t.Errorf("containsDuplicate(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, test := range tests {
		actual := isAnagram(test.s1, test.s2)
		if actual != test.expected {
			t.Errorf("isAnagram(%s, %s) = %v; expected %v", test.s1, test.s2, actual, test.expected)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, test := range tests {
		actual := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("twoSum(%v, %d) = %v; expected %v", test.nums, test.target, actual, test.expected)
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}
	for _, test := range tests {
		actual := groupAnagrams(test.strs)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("groupAnagrams(%v) = %v; expected %v", test.strs, actual, test.expected)
		}
	}
}

func TestTopKFrequest(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}
	for _, test := range tests {
		actual := topKFrequent(test.nums, test.k)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("topKFrequent(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1,2,3,4}, []int{24,12,8,6}},
		{[]int{-1,1,0,-3,3}, []int{0,0,9,0,0}},
	}
	for _, test := range tests {
		actual := productExceptSelf(test.nums)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("productExceptSelf(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
	}
	for _, test := range tests {
		actual := longestConsecutive(test.nums)
		if test.expected != actual {
			t.Errorf("longestConsecutive(%v) = %d; expected %d", test.nums, actual, test.expected)
		}
	}
}
