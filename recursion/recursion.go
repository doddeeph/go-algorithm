package recursion

// https://leetcode.com/problems/generate-parentheses/
func generateParentheses(n int) []string {
	result := []string{}
	var generate func(s string, open, close int)
	generate = func(s string, open, close int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}
		if open < n {
			generate(s+"(", open+1, close)
		}
		if close < open {
			generate(s+")", open, close+1)
		}
	}
	generate("", 0, 0)
	return result
}
