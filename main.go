package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./lexer"
)

func main() {
	filename := os.Args[1]

	currentDirectory, err := os.Getwd()

	if err != nil {
		log.Fatal("")
	}

	filepath := currentDirectory + "\\" + filename

	fileContents, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal("")
	}

	tokens := lexer.Start(fileContents)
	fmt.Println(tokens)
}