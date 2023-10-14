package main

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func divide(divident, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cant devide by zero! Didn't you attend primary school?")
	}
	return divident / divisor, nil
}

func main() {
	output, err := divide(10, 0)

	fmt.Println(output)
	if err != nil {
		fmt.Println(err)
	}

}
