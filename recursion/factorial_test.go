package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct{
		n int
		expected int
	}{
		{5, 120},
	}
	for _, test := range tests {
		actual := factorial(test.n)
		if test.expected != actual {
			t.Errorf("factorial(%d) = %d; expected %d", test.n, actual, test.expected)
		}
	}
}