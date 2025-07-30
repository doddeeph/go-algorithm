package backtracking

import "sort"

// https://leetcode.com/problems/subsets/description/
func subsets(nums []int) [][]int {
	var res [][]int
	var subsets func(start int, sub []int)
	subsets = func(start int, sub []int) {
		res = append(res, append([]int{}, sub...))
		for i := start; i < len(nums); i++ {
			sub = append(sub, nums[i])
			subsets(i+1, sub)
			sub = sub[:len(sub)-1]
		}
	}
	subsets(0, []int{})
	return res
}

// https://leetcode.com/problems/subsets-ii/description/
func subsetsWithDuplicate(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var subsets func(start int, sub []int)
	subsets = func(start int, sub []int) {
		res = append(res, append([]int{}, sub...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			sub = append(sub, nums[i])
			subsets(i+1, sub)
			sub = sub[:len(sub)-1]
		}
	}
	subsets(0, []int{})
	return res
}

// https://leetcode.com/problems/combination-sum/description/
func combinationSum(nums []int, target int) [][]int {
	var res [][]int
	var combinationSum func(start int, sum int, sub []int)
	combinationSum = func(start int, sum int, sub []int) {
		if sum == target {
			res = append(res, append([]int{}, sub...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(nums); i++ {
			sub = append(sub, nums[i])
			combinationSum(i, sum+nums[i], sub)
			sub = sub[:len(sub)-1]
		}
	}
	combinationSum(0, 0, []int{})
	return res
}

// https://leetcode.com/problems/combination-sum-ii/description/
func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var combinationSum func(start int, sum int, sub []int)
	combinationSum = func(start int, sum int, sub []int) {
		if sum == target {
			res = append(res, append([]int{}, sub...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			sub = append(sub, nums[i])
			combinationSum(i+1, sum+nums[i], sub)
			sub = sub[:len(sub)-1]
		}
	}
	combinationSum(0, 0, []int{})
	return res
}

// https://leetcode.com/problems/permutations/description/
func Permute(nums []int) [][]int {
	var res [][]int
	var permute func(sub []int, used []bool)
	permute = func(sub []int, used []bool) {
		if len(sub) == len(nums) {
			res = append(res, append([]int{}, sub...))
			return
		}
		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			sub = append(sub, num)
			permute(sub, used)
			sub = sub[:len(sub)-1]
			used[i] = false
		}
	}
	permute([]int{}, make([]bool, len(nums)))
	return res
}

// https://leetcode.com/problems/word-search/description/
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	var dfs func(r, c, start int) bool
	dfs = func(r, c, start int) bool {
		if start == len(word) {
			return true
		}
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[start] {
			return false
		}
		tmp := board[r][c]
		board[r][c] = '#'
		found := dfs(r+1, c, start+1) || dfs(r-1, c, start+1) || dfs(r, c+1, start+1) || dfs(r, c-1, start+1)
		board[r][c] = tmp
		return found
	}
	for r := range rows {
		for c := range cols {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}