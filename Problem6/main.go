// https://projecteuler.net/problem=6
// Sum square difference

// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of first ten natural numbers is (1+2+...+10)^2 = 55^2 = 3025
// difference = 3025 - 385 = 2640. Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

// 25164150

package main

import "fmt"

func main() {
	sumofSquares := 0
	squareofSum := 0
	difference := 0

	for i := 1; i <= 100; i++ {
		sumofSquares += i * i
		squareofSum += i
	}
	squareofSum *= squareofSum
	difference = squareofSum - sumofSquares
	fmt.Println(difference)
}
