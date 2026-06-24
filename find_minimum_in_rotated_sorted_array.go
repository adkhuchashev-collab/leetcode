package main

func findMin(nums []int) int {
	first, last := 0, len(nums)-1

	for first <= last {
		if first == last {
			return nums[first]
		}

		if nums[first] < nums[last] {
			return nums[first]
		}

		mid := (first + last) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[first] < nums[mid] {
			first = mid + 1
			continue
		}

		last = mid - 1
	}

	return 0
}
