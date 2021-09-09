package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type journal struct {
	entries []string
}

func (j *journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *journal) AddEntry(text string) int {

	entryCount++

	entry := fmt.Sprintf("%d: %s", entryCount, text)

	j.entries = append(j.entries, entry)

	return entryCount
}

// seperation of concerns

// bad approach
// func (j *journal) Save(filename string) {
// 	ioutil.WriteFile(filename, []byte(j.String()), 0644)
// }

// func (j *journal) Load(filename string) {
// 	//...
// }

// func (j *journal) LoadFromWeb(url *url.URL) {
// 	//...
// }

var LineSeperator = "\n"

func SaveToFile(j *journal, filename string) {
	ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeperator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *journal, filename string) {
	ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	myJournal := &journal{}
	myJournal.AddEntry("Hello World")
	myJournal.AddEntry("Learning about SRP")

	// save using a seperate function
	SaveToFile(myJournal, "journal.txt")

	// save using a method of a struct
	persist := &Persistence{lineSeparator: "\r\n"}
	persist.SaveToFile(myJournal, "journal2.txt")
}
