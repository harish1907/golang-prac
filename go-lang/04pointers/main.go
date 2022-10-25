package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the pointers.")

	var mynumber int = 20
	fmt.Println("mynumber",mynumber)

	ptr := &mynumber
	fmt.Println("ptr address:", ptr)
	*ptr += 2
	fmt.Println("ptr value:", *ptr)
}