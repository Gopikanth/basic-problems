package main

import "fmt"

func isPalindrome(n int) bool {
	r := 0
	o := n
	for n != 0 {
		r = (r * 10) + (n % 10)
		n = n / 10
	}
	return o == r

}
func main() {
	max := 999
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindrome(product) && product > max {
				max = product
			}
		}

	}
	fmt.Println(max)
}
	

