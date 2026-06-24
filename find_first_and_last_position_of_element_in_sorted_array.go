package main

func searchRange(nums []int, target int) []int {
	return []int{
		bound(nums, target, true),
		bound(nums, target, false),
	}
}

func bound(nums []int, target int, lower bool) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if lower {
				if mid == 0 || nums[mid] != nums[mid-1] {
					return mid
				}
				right = mid - 1

				continue

			}
			if mid == len(nums)-1 || nums[mid] != nums[mid+1] {
				return mid
			}
			left = mid + 1

			continue
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
