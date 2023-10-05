// File: pdfconverter/pdfconverter.go

package pdfconverter

import (
	"fmt"
	"os"
	"os/exec"
)

// ConvertPDFToText converts a PDF file to text.
func ConvertPDFToText(inputPDFPath, outputTextPath string) error {
	cmd := exec.Command("pdftotext", inputPDFPath, outputTextPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting PDF to text: %v", err)
	}

	return nil
}

// ConvertPDFToHTML converts a PDF file to HTML.
func ConvertPDFToHTML(inputPDFPath, outputHTMLPath string) error {
	cmd := exec.Command("pdftohtml", inputPDFPath, outputHTMLPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting PDF to HTML: %v", err)
	}

	return nil
}
