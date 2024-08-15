package main

func HitungFactorial(n int) int {
	var factorial int
	if n < 0 {
		factorial = -1
	} else if n == 0 {
		factorial = 1
	} else {
		factorial = n
		for i := 1; i < n; i++ {
			factorial *= (n - i)
		}
	}

	return factorial
}
