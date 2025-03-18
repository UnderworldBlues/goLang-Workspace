package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage: go run luhn.go <number>")
		os.Exit(1)
	}

	number := os.Args[1]

	if !isValid(number) {
		fmt.Println("invalid number")
		os.Exit(1)

	}

	if luhnAlgorithm(number) {
		fmt.Println("number is valid")
	} else {
		fmt.Println("number is invalid")
	}

}

func isValid(number string) bool {
	// iterate over each character
	for _, c := range number {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func luhnAlgorithm(number string) bool {

	n := len(number)
	sum := 0
	alternate := false

	for i := n - 1; i >= 0; i-- {

		digit, _ := strconv.Atoi(string(number[i]))

		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}
