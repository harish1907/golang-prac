package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Harish Chaudhary..")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
