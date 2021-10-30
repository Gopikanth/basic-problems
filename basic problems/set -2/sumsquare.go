package main

import "fmt"

func main() {
	max := 10
	sum_of_square := 0
	for i := 1; i <= max; i++ {
		sol := i * i
		sum_of_square = sum_of_square + sol

	}
	fmt.Println(sum_of_square)
}