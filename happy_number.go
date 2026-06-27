package main

import (
	"strconv"
)

func isHappy(n int) bool {
	hash := make(map[int]struct{})
	values := getArr(n)

	for {
		hash[n] = struct{}{}

		n = getSquareSum(values)

		values = getArr(n)
		if sum(values) == 1 {
			return true
		}

		if _, ok := hash[n]; ok {
			return false
		}
	}
}

func getSquareSum(values []int) int {
	var n int
	for _, digit := range values {
		n += digit * digit
	}
	return n
}

func sum(nums []int) int {
	var result int
	for _, v := range nums {
		result += v
	}
	return result
}

func getArr(n int) []int {
	var values []int
	str := strconv.Itoa(n)
	for _, v := range str {
		v = v - '0'
		values = append(values, int(v))
	}
	return values
}
