package main

import "fmt"

func main() {
	const b = 2021
	fmt.Println("Enter your year of birth")
	var a int
	fmt.Scanln(&a)
	age := b-a 
	fmt.Printf("your age  is  %v" ,age)
}