package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite int
	StreetAddress, City string
}

type Employee struct {
	Name string
	Office Address
}

var mainOfficeProto = Employee{
	Name:   "",
	Office: Address{
		Suite:         0,
		StreetAddress: "123 East Dr",
		City:          "London",
	},
}
var auxOfficeProto = Employee{
	Name:   "",
	Office: Address{
		Suite:         0,
		StreetAddress: "66 West Dr",
		City:          "London",
	},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func newMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOfficeProto, name, suite)
}
func newAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOfficeProto, name, suite)
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

func main()  {
	john := newMainOfficeEmployee("John", 100)
	jane := newAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
