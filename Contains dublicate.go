package main

func containsDuplicate(nums []int) bool {
	hash := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			return true
		}
		hash[v] = struct{}{}
	}
	return false
}
