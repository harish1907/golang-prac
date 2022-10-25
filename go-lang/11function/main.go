package main

import "fmt"

func main() {
	answer, message := add(12, 13)
	fmt.Printf("Answer : %v\nMessage: %v", answer, message)
	sliceadd := sliceadd("harish chaudhary", 10, 20, 30, 40, 50)
	fmt.Println("Slice add :", sliceadd)
}

func add(val1 int, val2 int) (int, string) {
	return val1 + val2, "Hello there i am add two value"
}

func sliceadd(name string, values ...int) int {
	fmt.Println("name:", name)
	total := 0
	for i := range values {
		total += i
	}
	return total
}
