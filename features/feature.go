// File: pdfconverter/features/feature.go

package features

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NewFeature() {
	fmt.Println("This is a new feature!")
}

// PDFMetadata struct to store PDF metadata
type PDFMetadata struct {
	Title       string
	Author      string
	CreationDate string
}



// WordCount counts the number of words in the specified text file.
func WordCount(textFilePath string) (int, error) {
	file, err := os.Open(textFilePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return wordCount, nil
}





// GetPDFMetadata retrieves metadata from a PDF file
func GetPDFMetadata(inputPDFPath string) (*PDFMetadata, error) {
	cmd := exec.Command("pdfinfo", inputPDFPath)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting PDF metadata: %v", err)
	}

	// Parse the output to extract metadata
	pdfMetadata := &PDFMetadata{}
	fmt.Sscanf(string(output), "Title: %[^\n]\nAuthor: %[^\n]\nCreationDate: %[^\n]", &pdfMetadata.Title, &pdfMetadata.Author, &pdfMetadata.CreationDate)

	return pdfMetadata, nil
}
