// https://projecteuler.net/problem=5
// https://en.wikipedia.org/wiki/Smallest_multiple

// 2520 is the smallest number that can be divided by
// each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all numbers from 1 to 20?

package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		isDivisible := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			fmt.Println(i)
			break
		}
	}
}
