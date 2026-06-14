package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	start, end := 0, n
	var lastKnownBad int

	for start <= end {

		mid := (start + end) / 2

		if !isBadVersion(mid) {
			start = mid + 1
			continue
		}

		if start == mid {
			return mid
		}

		lastKnownBad = mid
		end = mid - 1
	}

	return lastKnownBad
}

// given by default in leetCode
func isBadVersion(n int) bool {
	return true
}
