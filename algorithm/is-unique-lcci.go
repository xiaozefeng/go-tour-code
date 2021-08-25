package algorithm

import "strings"

func isUnique(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k,v := range s {
		if v >127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
