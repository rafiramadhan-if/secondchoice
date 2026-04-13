package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Deret Fibonacci (suku ke-0 hingga ke-10):")
	
	fmt.Print("n  | ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%-4d", i)
	}
	
	fmt.Println()
	fmt.Print("Sn | ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%-4d", fibonacci(i))
	}
	fmt.Println()
}
