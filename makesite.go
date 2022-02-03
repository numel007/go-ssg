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

// type Page struct {
// 	TextFilePath string
// 	TextFileName string
// 	HTMLFilePath string
// 	Content      string
// }

func main() {
	// inputFile := flag.String("file", "first-post.txt", "txt file to pass in")
	inputDir := flag.String("dir", " ", "Return all .txts in directory")
	flag.Parse()

	if *inputDir != " " {
		println("Converting all txt files in " + *inputDir + "/")
		files, fileError := ioutil.ReadDir(*inputDir)
		check(fileError)

		for _, file := range files {
			if file.IsDir() {
			} else {
				fileNameArray := strings.Split(file.Name(), ".")
				if fileNameArray[len(fileNameArray)-1] == "txt" {
					outputName := strings.Split(file.Name(), ".")[0] + ".html"
					data, _ := ioutil.ReadFile(file.Name())

					myStruct := Bullshit{Data: string(data)}

					// Use a defined template
					parsedTemplate, _ := template.ParseFiles("template.tmpl")

					// Create a file to write to
					newFile, _ := os.Create(outputName)

					// Write to new file using template and data
					err := parsedTemplate.Execute(newFile, myStruct)
					check(err)
				}
			}
		}
	} else {
		println("-dir not passed. No htmls generated.")
	}
}
