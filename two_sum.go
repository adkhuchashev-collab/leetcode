package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	hash := make(map[int][]int, len(nums))
	for i, num := range nums {
		hash[num] = append(hash[num], i)
	}

	for digit, guess1 := range hash {
		if guess2, ok := hash[target-digit]; ok {
			if guess1[0] == guess2[0] {
				if len(guess1) < 2 {
					continue
				}
				return guess1[:2]
			}
			return []int{guess1[0], guess2[0]}
		}
	}

	return []int{}
}
