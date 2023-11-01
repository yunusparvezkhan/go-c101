package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizzbuzz,\n")
		}

		if i%3 == 0 {
			fmt.Printf("fizz,\n")
		} else if i%5 == 0 {
			fmt.Printf("buzz,\n")
		} else {
			fmt.Printf(fmt.Sprintf("%d,\n", i))
		}
	}

}

// don't touch below this line

func main() {
	fizzbuzz()
}
