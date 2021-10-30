package main

import "fmt"

func main() {
	fmt.Println("Enter the weights")
	var a int 
	fmt.Println("Enter the 1st weight")
	fmt.Scanln(&a)
	var b int
	fmt.Println("Enter the 2nd weight")
	fmt.Scanln(&b)
	var c int
	fmt.Println("Enter the 3rd weight")
	fmt.Scanln(&c)
	sum := a+b+c
	avg := sum /3
	fmt.Printf("The average of weights are %v",avg)
	
}