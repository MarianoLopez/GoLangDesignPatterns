package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

func main() {
	john := Person{
		Name:    "john",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "312 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	/*jane := john
	jane.Name = "Jane"
	jane.Address = &Address{
		StreetAddress: john.Address.StreetAddress,
		City:          john.Address.City,
		Country:       john.Address.Country,
	}
	jane.Address.StreetAddress = "321 Baker St"*/

	/*jane := john
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"*/

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}