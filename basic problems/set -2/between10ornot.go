//to find the input is between 10 or not
package main

import "fmt"

func main() {
	fmt.Println("Enter the Number")
	var a int
	fmt.Scanln(&a)
	if a <= 10 {
		fmt.Println("the number is in between 10")
	}else{
		fmt.Println("the number is not between 10")
	}
}
