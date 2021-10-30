package main

import "fmt"

func main() {
	fmt.Println("Enter your name")
	var a  string
	fmt.Scanln(&a)
	fmt.Printf("your name is %s",a)
}