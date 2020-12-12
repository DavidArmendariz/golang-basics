package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{
		firstName: "David",
		lastName:  "Armendariz",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 123456,
		},
	}

	p.updateName("Adrian")
	p.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
