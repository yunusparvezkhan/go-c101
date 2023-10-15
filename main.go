package main

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func divide(divident, divisor float32) (float32, string) {
	if divisor == 0 {
		return 0, "cant devide by zero! Didn't you attend primary school?"
	}
	return divident / divisor, ""
}

func main() {
	output, err := divide(10, 0)

	if err != "" {
		fmt.Println("Error dividing values:", err)
	} else {
		fmt.Println(output)
	}

}
