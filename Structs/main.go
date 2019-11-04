package main

import "fmt"

type contactInfo struct {
	email    string
	postCode string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:    "jim@gmail.com",
			postCode: "SW2 2JR",
		},
	}

	fmt.Printf("%+v", jim)
}
