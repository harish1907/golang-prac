package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go lang.")
	furitList := []string{"apple", "banana", "mango"}
	fmt.Println(furitList)

	furitList = append(furitList, "kiwi")
	fmt.Println(furitList)

	furitList = append(furitList, "guawa", "orange")
	fmt.Println(furitList)

	highscores := make([]int, 4)

	highscores[0] = 324
	highscores[1] = 768
	highscores[2] = 876
	highscores[3] = 432

	highscores = append(highscores, 890, 777)

	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	courses := []string{"Reactjs", "python", "ruby", "javascript", "golang"}
	index := 1
	fmt.Println(courses)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
