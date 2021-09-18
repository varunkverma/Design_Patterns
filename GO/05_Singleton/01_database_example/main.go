package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once init() -- thread safety
// laziness

var once sync.Once // ensures that something gets called only once

var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	// using once to do the passed func only once
	once.Do(func() {
		caps, err := readData("capitals.txt")
		db := &singletonDatabase{
			capitals: caps,
		}
		if err == nil {
			db.capitals = caps
		}
		instance = db
	})
	return instance
}

//utility func
func readData(path string) (map[string]int, error) {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)
	// fmt.Println("exPath: ", exPath)
	file, err := os.Open(path)
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

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println(pop)
}
