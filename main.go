package main

import (
	"fmt"
	"htmlLink/link"
)

func main() {

	// The HTML file to Parse
	const fileName = "example2.html"

	// GetLinks will take HTML 'fileName' and returns all the links in HTML page in the Link data structure format.
	links, err := link.GetLinks(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(links)
}
