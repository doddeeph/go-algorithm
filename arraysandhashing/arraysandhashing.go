package arraysandhashing

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s1 {
		count[ch]++
	}
	for _, ch := range s2 {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		res := target - num
		if idx, val := m[res]; val {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{}
}

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		m[sorted] = append(m[sorted], s)
	}
	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}

// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int, len(nums))
	for _, num := range nums {
		freqMap[num]++
	}
	type freqPair struct {
		num  int
		freq int
	}
	freqPairs := make([]freqPair, 0, len(freqMap))
	for num, freq := range freqMap {
		freqPairs = append(freqPairs, freqPair{num, freq})
	}
	sort.Slice(freqPairs, func(i, j int) bool {
		return freqPairs[i].freq > freqPairs[j].freq
	})
	result := make([]int, 0, k)
	for i := range k {
		result = append(result, freqPairs[i].num)
	}
	return result
}

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 1
	for i := range nums {
		result[i] = left
		left *= nums[i]
	}

	right := 1
	for i := n-1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	return true
}

// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	maxLength := 0
	for _, num := range nums {
		if !numSet[num-1] {
			currNum := num
			length := 1
			for numSet[currNum+1] {
				currNum++
				length++
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

// EY Careers
// doddee.ph@gmail.com
// T-@8dS79QX^PGt,
