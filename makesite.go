package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Bullshit struct {
	Data string
}

func createHTML(data string, outputName string) {
	newStruct := Bullshit{Data: data}
	parsedTemplate, _ := template.ParseFiles("template.tmpl")
	newFile, _ := os.Create(outputName)
	err := parsedTemplate.Execute(newFile, newStruct)
	check(err)
}

func main() {
	// inputFile := flag.String("file", "first-post.txt", "txt file to pass in")
	inputDir := flag.String("dir", " ", "Return all .txts in directory")
	flag.Parse()

	if *inputDir != " " {
		println("Converting all .txt & .md files in " + *inputDir + "/")
		files, fileError := ioutil.ReadDir(*inputDir)
		check(fileError)

		for _, file := range files {
			if file.Mode().IsRegular() {
				fileNameArray := strings.Split(file.Name(), ".")
				outputName := strings.Split(file.Name(), ".")[0] + ".html"

				if fileNameArray[len(fileNameArray)-1] == "txt" {
					data, _ := ioutil.ReadFile(file.Name())

					createHTML(string(data), outputName)

				} else if fileNameArray[len(fileNameArray)-1] == "md" {
					mdData, _ := ioutil.ReadFile(file.Name())
					mdDataConverted := string(markdown.ToHTML(mdData, nil, nil))

					createHTML(mdDataConverted, outputName)
				}
			}
		}
	} else {
		println("-dir not passed. No htmls generated.")
	}
}
