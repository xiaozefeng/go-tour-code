package algorithm

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var m = map[byte]int{}
	var ans = 0
	for i, j := 0, 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok {
			j = max(j, val+1)
		}
		m[s[i]] = i
		ans = max(ans, i-j+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
