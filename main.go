package main

import (
	"io"
	"log"
	"os"

	"github.com/mvenkatesh431/LinkParser/link"
)

func main() {

	// The HTML file to Parse
	const fileName = "example4.html"

	// ReadHTML will take fileName and returns the io.Reader
	reader, err := readHTML(fileName)
	if err != nil {
		log.Fatalf("Failed to read the HTML File %v", err)
	}

	// GetLinks will take 'io.Reader' and returns all the links in HTML page in the Link data structure format.
	links, err := link.GetLinks(reader)
	if err != nil {
		log.Fatalf("Failed to get links %v", err)
	}
	log.Println(links)
}

// readHTML will take fileName and returns the io.Reader
func readHTML(fileName string) (io.Reader, error) {
	// os.Open returns the io.Reader
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
