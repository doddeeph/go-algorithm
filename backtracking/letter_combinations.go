package backtracking

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	digitToLetter := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var solution func(idx int, sub string)
	solution = func(idx int, sub string) {
		if idx == len(digits) {
			res = append(res, sub)
			return
		}
		letter := digitToLetter[digits[idx]]
		for i := 0; i < len(letter); i++ {
			solution(idx+1, sub+string(letter[i]))
		}
	}
	solution(0, "")
	return res
}
