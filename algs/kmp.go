package main

func genNext(pattern string) []int {
	patternLen := len(pattern)
	j, next := 0, make([]int, patternLen)
	for i := 2; i < patternLen; i++ {
		for j != 0 && pattern[j] != pattern[i-1] {
			j = next[j]
		}
		if pattern[j] == pattern[i-1] {
			j++
		}
		next[i] = j
	}
	return next
}

func kmp(main, pattern string) int {
	mainLen, patternLen := len(main), len(pattern)
	j, next := 0, genNext(pattern)
	for i := 0; i < mainLen; i++ {
		for j > 0 && main[i] != pattern[j] {
			j = next[j]
		}
		if main[i] == pattern[j] {
			j++
		}
		if j == patternLen {
			return i - patternLen + 1
		}
	}
	return -1
}

func main() {
	println(kmp("GTGTGAGCTGGTGTGTGCFAA", "GTGTGCF"))
}
