package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	var username string = "harish chaudhary"
	fmt.Println(username, reflect.TypeOf(username))

	var price = 500.54
	fmt.Println(price, reflect.TypeOf(price))

	// walrus operator only use in function.
	isloging := true
	fmt.Println(isloging, reflect.TypeOf(isloging))

	var money = 30000
	fmt.Println("total money i have:", money, reflect.TypeOf(money))

	// User input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter th rating of our pizza:")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanx for give us rating, :", input)
	fmt.Printf("Type of variable %T", input)

	
	fmt.Println("Enter the user input value:")
	userinput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	print(userinput)
}
