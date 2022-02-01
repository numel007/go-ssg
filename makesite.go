package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
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
	inputFile := flag.String("file", "first-post.txt", "txt file to pass in")
	flag.Parse()

	outputName := strings.Split(*inputFile, ".")[0] + ".html"
	data, _ := ioutil.ReadFile(*inputFile)

	myStruct := Bullshit{Data: string(data)}

	// Use a defined template
	parsedTemplate, _ := template.ParseFiles("template.tmpl")

	// Create a file to write to
	newFile, _ := os.Create(outputName)

	// Write to new file using template and data
	err := parsedTemplate.Execute(newFile, myStruct)
	check(err)
}
