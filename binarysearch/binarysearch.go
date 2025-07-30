package binarysearch

// https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return false
	}
	l, r := 0, m*n-1
	for l <= r {
		mid := l + (r-l)/2
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		} else if val < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMinRotatedSortedArray(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func searchRotatedSortedArray(nums []int, target int) int {
	return -1
}

