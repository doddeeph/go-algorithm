package recursion

// Time Complexity: O(2^n)
// Space Complexity: O(1)
func fibonacci(n int) int {
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return fibonacci(n)
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func fib(n int) int {
	var fib func(n int) int
	fib = func(n int) int {
		dp := make([]int, n+1)
		dp[0], dp[1] = 0, 1
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
	}
	return fib(n)
}


// Time Complexity: O(n)
// Space Complexity: O(1)
func fibSpaceOptimized(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}