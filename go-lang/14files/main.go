package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file handing in golang.")
	content := "Hello there this is first time i learn golang language.."

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	defer file.Close()
}
