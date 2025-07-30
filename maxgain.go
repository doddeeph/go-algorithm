package main

import "sort"

func maxGain(nums []int, m int) int {
	sort.Ints(nums)
	res := 0;
	for _, num := range nums {
		if num < 0 && m > 0 {
			res += num * -1
			m--
		}
	}
	return res
}