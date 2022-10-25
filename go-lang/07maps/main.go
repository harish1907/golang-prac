package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang...")
	courses := make(map[string]string)

	courses["js"] = "javascript"
	courses["py"] = "python"
	courses["rb"] = "ruby"
	courses["go"] = "golang"

	fmt.Println(courses)
	fmt.Println("javascript short: ", courses["js"])

	delete(courses, "py")
	fmt.Println(courses)

	for key, value := range courses {
		fmt.Printf("For key %v, For value %v\n", key, value)
	}
}
