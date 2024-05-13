func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashTableS := make(map[byte]int)
	hashTableT := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := hashTableS[s[i]]; ok {
			hashTableS[s[i]]++
			continue
		}
		hashTableS[s[i]] = 1
	}
	for j := 0; j < len(t); j++ {
		if _, ok := hashTableT[t[j]]; ok {
			hashTableT[t[j]]++
			continue
		}
		hashTableT[t[j]] = 1
	}
	for v, n := range hashTableS {
		if hashTableT[v] != n {
			return false
		}
	}
	return true
}