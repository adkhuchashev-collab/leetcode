package main

func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		hash[v] = struct{}{}
	}

	var result []int
	for _, v := range nums2 {
		if _, ok := hash[v]; ok {
			result = append(result, v)
			delete(hash, v)
		}
	}

	return result
}
