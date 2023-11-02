package main

import (
	"fmt"
)

func printPrimes(max int) {
	// Starting with 2, because 1 is not a prime number
	for i := 2; i < max; i++ {
		if i == 2 {
			fmt.Println(i)
		}
		if i%2 == 0 {
			continue
		}
		if (i*i)%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
