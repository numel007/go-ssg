package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Bullshit struct {
	Data string
}

func main() {
	data, _ := ioutil.ReadFile("./first-post.txt")

	myStruct := Bullshit{Data: string(data)}

	// Use a defined template
	parsedTemplate, _ := template.ParseFiles("template.html")

	// Create a file to write to
	newFile, _ := os.Create("first-post.html")

	// Write to new file using template and data
	err := parsedTemplate.Execute(newFile, myStruct)
	check(err)
}
