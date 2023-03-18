package main

import (
	"fmt"
	"io/ioutil"
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	converter := md.NewConverter("", true, nil)

	// read file `assets/article.html`

	filePath := "assets/article.html"
	fileData, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	html := string(fileData)

	markdown, errReadingFile := converter.ConvertString(html)
	if errReadingFile != nil {
		log.Fatal(errReadingFile)
	}

	// write markdown to file `assets/article.md`

	filePath = "assets/article.md"

	errorWritingFile := ioutil.WriteFile(filePath, []byte(markdown), 0644)

	if errorWritingFile != nil {
		log.Fatal(errorWritingFile)
		return
	}

	fmt.Println("Done!")
}
