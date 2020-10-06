package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type JournalManagement interface {
	AddEntry(text string) int
	String() string
}

type Journal struct {
	entries []string
}

type DefaultJournalManagement struct {
	journal *Journal
}

func (journalManagement *DefaultJournalManagement) String() string {
	return strings.Join(journalManagement.journal.entries, "\n")
}

func (journalManagement *DefaultJournalManagement) AddEntry(text string) int {
	entryCount := len(journalManagement.journal.entries) + 1
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	journalManagement.journal.entries = append(journalManagement.journal.entries, entry)
	return entryCount
}

//BREAK PRINCIPLE
func (journalManagement *DefaultJournalManagement) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(journalManagement.String()), 0644)
}

//PERSISTENCE should be handled by another struct e.g
type JournalFileRepository interface {
	saveToFile(j *Journal, filename string)
}
type JournalFileRepositoryImpl struct {
	lineSeparator string
}

func (p *JournalFileRepositoryImpl) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	journal := &Journal{}

	var journalManagement JournalManagement = &DefaultJournalManagement{
		journal: journal,
	}
	journalManagement.AddEntry("Learning Golang")
	journalManagement.AddEntry("Java please come back(?")
	fmt.Print(journalManagement.String())

	var journalFileRepository JournalFileRepository = &JournalFileRepositoryImpl{lineSeparator: "\n"}
	journalFileRepository.saveToFile(journal, "journals.txt")
}
