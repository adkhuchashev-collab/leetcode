package main

func fib(n int) int {
	if n <= 1 {
		return n
	}

	//last 2 numbers
	a, b := 0, 1

	//starts from second number
	for i := 2; i < n; i++ {
		c := a + b
		a, b = b, c
	}

	return b
}
