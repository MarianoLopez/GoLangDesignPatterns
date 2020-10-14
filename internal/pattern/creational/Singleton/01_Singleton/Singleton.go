package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

//read capitals.txt
func readData(path string) (map[string]int, error) {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

//options to avoid concurrency problems
//sync.Once init() - thread safety
//laziness  (right now)
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("/DesignPatterns/internal/pattern/creational/Singleton/01_Singleton/capitals.txt")
		if err != nil {
			fmt.Print(err)
		}
		db := singletonDatabase{capitals: caps}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul = ", pop)
}