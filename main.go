package main

import "fmt"

func main() {
	fmt.Println(findPrimesInRange(5, 78))
	fmt.Println(findPrimesInRange(78, 5))
}

// Является ли число простым
func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimesInRange(min, max int) []int {
	if min > max {
		min, max = max, min
	}
	var primes []int
	for i := min; i <= max; i++ {
		if Prime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}