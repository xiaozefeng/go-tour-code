package algorithm

// 1. sort and compare
// 2. stat s and t in map
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m = make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}
	for _, ch := range t {
		m[ch]--
		if m[ch] <0 {
			return false
		}
	}
	return true
}

// solution 3
//func isAnagram(s string, t string) bool {
//	var t1,t2 [26]int
//	n1:= len(s)
//	n2:= len(t)
//	if n1!=n2{
//		return false
//	}
//	for _, ch := range s {
//		t1[ch-'a']++
//	}
//	for _, ch := range t {
//		t2[ch-'a']++
//	}
//	return t1 == t2
//}

// solution 2
//func isAnagram(s string, t string) bool {
//	var table [26]int
//
//	n1:= len(s)
//	n2:= len(t)
//	if n1!=n2{
//		return false
//	}
//	for i := 0; i < n1; i++ {
//		table[s[i] - 'a'] ++
//		table[t[i] -'a'] --
//	}
//	for i := 0; i < len(table); i++ {
//		if table[i] < 0 {
//			return false
//		}
//	}
//	return true
//}

// solution 1
//func isAnagram(s string, t string) bool {
//	s1 := []byte(s)
//	s2 := []byte(t)
//	sort.Slice(s1, func(i, j int) bool {
//		return s1[i] < s1[j]
//	})
//	sort.Slice(s2, func(i, j int) bool {
//		return s2[i] < s2[j]
//	})
//	return string(s1) == string(s2)
//}
