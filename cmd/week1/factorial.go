package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	n := 10

	fmt.Printf("Factorial de %d: %d\n", n, factorial(n))
}
