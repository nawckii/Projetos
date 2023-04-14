package main

func Fatorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * Fatorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
