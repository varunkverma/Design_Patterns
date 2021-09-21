package main

import (
	"fmt"
	"strings"
)

// type User struct {
// 	FullName string
// }

// func NewUser(fullName string) *User {
// 	return &User{
// 		FullName: fullName,
// 	}
// }

var allNames []string

type User struct {
	names []uint8 // flyweight
}

func NewUser(fullName string) *User {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := &User{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return result
}

func (u *User) FullName() string {
	parts := make([]string, 0)
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {
	john := NewUser("John Doe")
	jane := NewUser("Jane Doe")
	alsoJane := NewUser("Jane smith")

	// fmt.Println(john.FullName)
	// fmt.Println(jane.FullName)
	// fmt.Println(alsoJane.FullName)

	// fmt.Println("Memory usable")

	// fmt.Println("Total memory used:",
	// 	len([]byte(john.FullName))+
	// 		len([]byte(jane.FullName))+
	// 		len([]byte(alsoJane.FullName)), "bytes.")

	fmt.Println(john.FullName())
	fmt.Println(jane.FullName())
	fmt.Println(alsoJane.FullName())

	fmt.Println("Memory usable")

	totalMemory := 0
	for _, a := range allNames {
		totalMemory += len([]byte(a))
	}
	totalMemory += len(john.names)
	totalMemory += len(jane.names)
	totalMemory += len(alsoJane.names)

	fmt.Println("Total memory used:", totalMemory, "bytes.")
}
