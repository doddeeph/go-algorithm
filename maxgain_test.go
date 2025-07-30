package main

import "testing"

func TestMaxGain(t *testing.T) {
	tests := []struct {
		nums     []int
		m        int
		expected int
	}{
		{[]int{-1, -10}, 1, 10},
		{[]int{10, 20, 30}, 2, 0},
	}
	for _, test := range tests {
		actual := maxGain(test.nums, test.m)
		if test.expected != actual {
			t.Errorf("maxGain(%v, %d) = %d; expected %d", test.nums, test.m, actual, test.expected)
		}
	}
}
