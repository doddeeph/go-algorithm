package recursion

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct{
		n int
		expected int
	}{
		{5, 5},
	}
	for _, test := range tests {
		actual := fibonacci(test.n)
		if test.expected != actual {
			t.Errorf("fibonacci(%d) = %d; expected %d", test.n, actual, test.expected)
		}
	}
}

func TestFib(t *testing.T) {
	tests := []struct{
		n int
		expected int
	}{
		{5, 5},
	}
	for _, test := range tests {
		actual := fib(test.n)
		if test.expected != actual {
			t.Errorf("fib(%d) = %d; expected %d", test.n, actual, test.expected)
		}
	}
}

func TestFibSpaceOptimized(t *testing.T) {
	tests := []struct{
		n int
		expected int
	}{
		{5, 5},
	}
	for _, test := range tests {
		actual := fibSpaceOptimized(test.n)
		if test.expected != actual {
			t.Errorf("fib(%d) = %d; expected %d", test.n, actual, test.expected)
		}
	}
}