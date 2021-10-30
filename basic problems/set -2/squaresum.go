package main

import "fmt"

func main() {
	max := 10
	square_of_sum := 0
	for i := 1; i <= max; i++ {
		square_of_sum = square_of_sum + i

	}
	fmt.Println(square_of_sum)
}
