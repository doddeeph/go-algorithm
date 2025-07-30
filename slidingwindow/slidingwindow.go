package slidingwindow

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start, maxLen := 0, 0
	for end, ch := range s {
		if i, ok := m[ch]; ok && i >= start {
			start = i + 1
		}
		m[ch] = end
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}
	return maxLen
}

// https://leetcode.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	count := [26]int{}
	left, maxCount, maxLen := 0, 0, 0
	for right, _ := range s {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxCount {
			maxCount = count[s[right]-'A']
		}
		if right-left+1-maxCount > k {
			count[s[left]-'A']--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right-left+1
		}
	}
	return maxLen
}
