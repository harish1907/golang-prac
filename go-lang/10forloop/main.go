package main

import "fmt"

func main() {
	fmt.Println("For loop in golang.")

	weeklist := []string{"sunday", "monday", "tuesday", "wednesday", "saturday"}
	fmt.Println(weeklist)

	// for i := 0; i < len(weeklist); i++ {
	// 	fmt.Println(weeklist[i])
	// }

	// for i := range weeklist {
	// 	fmt.Println(weeklist[i])
	// }

	// for index, value := range weeklist {
	// 	fmt.Printf("Index is %v and value is %v\n", index, value)
	// }

	rougevalue := 1

	for rougevalue < 10 {

		if rougevalue == 7 {
			goto ico
		}

		if rougevalue == 5 {
			rougevalue++
			continue
		}

		fmt.Println("Value is", rougevalue)
		rougevalue++
	}

ico:
	fmt.Println("Value is 7 and then run goto..")
}
