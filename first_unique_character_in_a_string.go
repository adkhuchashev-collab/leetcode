package main

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	hash := make(map[rune]int, len([]rune(s)))
	for _, v := range s {
		hash[v]++
	}

	for i, v := range s {
		if q, ok := hash[v]; ok {
			if q == 1 {
				return i
			}
		}
	}

	return -1
}
