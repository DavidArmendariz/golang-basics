package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	david := person{
		firstName: "David",
		lastName:  "Armendariz",
		contact: contactInfo{
			email:   "test@test.com",
			zipCode: 123456,
		},
	}
	fmt.Printf("%+v", david)
}
