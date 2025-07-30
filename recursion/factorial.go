package recursion

func factorial(n int) int {
	var factorial func(n int) int
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}
	return factorial(n)
}
