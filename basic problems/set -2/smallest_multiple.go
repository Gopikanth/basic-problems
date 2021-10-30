package main

import "fmt"

func main() {
	max := 2520 //using as max because 2520 is divisible by 1 to 10
	for i := max; ; i++ {
		can := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				can = false
				break
			}
		}
		if can {
			max = i
			break
		}
	}
	fmt.Println(max)
}
