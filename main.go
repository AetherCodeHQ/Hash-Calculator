package main

import (
	"fmt"
	"os"
)

// hash_calculator - Calculate file hashes
func hash_calculator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Hash-Calculator")
	fmt.Println("  Calculate file hashes")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	hash_calculator(path)
}
