package main

import "fmt"

// generateFibonacci generates a Fibonacci sequence up to n terms
func generateFibonacci(n int) []int {
    fib := make([]int, n)
    if n >= 1 {
        fib[0] = 0
    }
    if n >= 2 {
        fib[1] = 1
    }
    for i := 2; i < n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    return fib
}

func main() {
    fmt.Println(generateFibonacci(10))
}
