package main

func isAnagram(s string, t string) bool {
	hash := make(map[rune]int, len([]rune(s)))
	for _, letter := range s {
		hash[letter]++
	}

	for _, letter := range t {
		if _, ok := hash[letter]; ok {
			hash[letter]--
			if hash[letter] == 0 {
				delete(hash, letter)
			}
		} else {
			return false
		}
	}

	if len(hash) == 0 {
		return true
	}
	return false
}
