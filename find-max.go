package main

import "fmt"

// findMax finds the maximum value in a slice of integers
func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0 // or you could return an error
    }
    max := numbers[0]
    for _, num := range numbers[1:] {
        if num > max {
            max = num
        }
    }
    return max
}

func main() {
    fmt.Println(findMax([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}))
}
