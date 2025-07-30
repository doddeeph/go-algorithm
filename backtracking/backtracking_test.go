package backtracking

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for _, test := range tests {
		actual := subsets(test.nums)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("subsets(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}

func TestSubsetsWithDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for _, test := range tests {
		actual := subsetsWithDuplicate(test.nums)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("subsetsWithDuplicate(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		// {[]int{2}, 1, [][]int{}},
	}
	for _, test := range tests {
		actual := combinationSum(test.nums, test.target)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("combinationSum(%v, %d) = %v; expected %v", test.nums, test.target, actual, test.expected)
		}
	}
}

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}
	for _, test := range tests {
		actual := combinationSum2(test.nums, test.target)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("combinationSum2(%v, %d) = %v; expected %v", test.nums, test.target, actual, test.expected)
		}
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, test := range tests {
		actual := Permute(test.nums)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("permute(%v) = %v; expected %v", test.nums, actual, test.expected)
		}
	}
}

func TestExist(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
	}
	for _, test := range tests {
		actual := exist(test.board, test.word)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("exist(%v, %s) = %v; expected %v", test.board, test.word, actual, test.expected)
		}
	}
}
