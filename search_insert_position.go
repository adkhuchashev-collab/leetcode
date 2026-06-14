package main

func searchInsert(nums []int, target int) int {

	start, end := 0, len(nums)-1

	for start <= end {

		mid := (start + end) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target && start != end:
			end = mid - 1
		case nums[mid] < target && start != end:
			start = mid + 1
		case start == end:
			if nums[mid] > target {
				return mid - 1
			}
			return mid + 1
		}
	}

	return -1
}
