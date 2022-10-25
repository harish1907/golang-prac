package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main()  {
	fmt.Println("Welcome to our pizza app.")
	fmt.Println("Give us a rating in between 1 to 5.")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	stringconvert, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(stringconvert+1)
	}
}