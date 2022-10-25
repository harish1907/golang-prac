package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang..")
	//There is no inheritance and there is no classes they have structs.
	harish := User{"harish", "harishchaudhary920@gmail.com", true, 22}
	fmt.Println("Complete strut details:", harish)
	fmt.Printf("Struct with all key values: %+v\n", harish)

	fmt.Printf("name of user %v\n email of user %v\n status of user %v\n age of user %v\n", harish.Name, harish.Email, harish.Status, harish.Age)
	fmt.Println(harish.checkStatus())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) checkStatus() string {
	if u.Status {
		return "This user is logged in."
	} else {
		return "This user is logged out"
	}
}
