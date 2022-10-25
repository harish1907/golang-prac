package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang...")
	var frutList [4]string

	frutList[0] = "Apple"
	frutList[1] = "tomato"
	frutList[2] = ""
	frutList[3] = "Apple"
	fmt.Println(frutList)
}
