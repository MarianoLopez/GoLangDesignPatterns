package main

import (
	"fmt"
)

type Database interface {
	GetPopulation(name string) int
}

func GetTotalPopulation(db Database, cities []string) int {
	result := 0

	for _, city := range cities {
		result += db.GetPopulation(city) //no DI -> GetSingletonDatabase().GetPopulation
	}

	return result
}

type DummyDatabase struct {
	dummyData map[string] int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha" : 1,
			"beta"  : 2,
			"gamma" : 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulation(&DummyDatabase{}, names)
	fmt.Println(tp)
}