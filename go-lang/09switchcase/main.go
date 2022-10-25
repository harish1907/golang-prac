package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang..")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Dice number is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice number is 1 you can open.")
	case 2:
		fmt.Println("Dice number 2 move 2 spot.")
	case 3:
		fmt.Println("Dice number 3 move 3 spot.")
		fallthrough
	case 4:
		fmt.Println("Dice number 4 move 4 spot.")
	case 5:
		fmt.Println("Dice number 5 move 5 spot.")
	case 6:
		fmt.Println("Dice number 6 move 6 spot and you can roll again..")
	default:
		fmt.Println("What was that?")
	}
}
