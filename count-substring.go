package main

import (
    "fmt"
    "strings"
)

// countSubstring counts the occurrences of a substring in a string
func countSubstring(s, substr string) int {
    return strings.Count(s, substr)
}

func main() {
    fmt.Println(countSubstring("Go is awesome. Go is fast.", "Go"))
}
