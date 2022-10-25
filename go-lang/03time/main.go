package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study..")
	presentTime := time.Now()
	fmt.Println("presentTime", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 3:04:05 PM Monday"))

	createdAt := time.Date(2001, time.March, 23, 10, 10, 20, 30, time.Local)
	fmt.Println("createdAt", createdAt)
	fmt.Println(createdAt.Format("01-02-2006 Monday"))
}
