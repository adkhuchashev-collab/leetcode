package main

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)*len(matrix[0])-1

	for l <= r {
		mid := (l + r) / 2
		row := mid / len(matrix[0])
		column := mid % len(matrix[0])

		if matrix[row][column] == target {
			return true
		}

		if matrix[row][column] < target {
			l = mid + 1
		}

		if matrix[row][column] > target {
			r = mid - 1
		}
	}

	return false
}
