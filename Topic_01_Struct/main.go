package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// Save a page as a file with title = name file
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// Construct a new Page object from the file
// return a pointer to the Page object in memory
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, err
}

func main1() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, err := loadPage("TestPage")
	if err != nil {
		fmt.Println("Load file failed")
	} else {
		fmt.Println(string(p2.Body))
	}
}
