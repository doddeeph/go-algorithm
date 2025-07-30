package binarysearch

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for _, test := range tests {
		actual := search(test.nums, test.target)
		if test.expected != actual {
			t.Errorf("search(%v, %d) = %d; expected %d", test.nums, test.target, actual, test.expected)
		}
	}
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}
	for _, test := range tests {
		actual := searchMatrix(test.matrix, test.target)
		if test.expected != actual {
			t.Errorf("searchMatrix(%v, %d) = %v; expected %v", test.matrix, test.target, actual, test.expected)
		}
	}
}

func TestFindMinRotatedSortedArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}
	for _, test := range tests {
		actual := findMinRotatedSortedArray(test.nums)
		if test.expected != actual {
			t.Errorf("findMin(%v) = %d; expected %d", test.nums, actual, test.expected)
		}
	}
}
