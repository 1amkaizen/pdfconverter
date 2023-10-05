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

// GetPDFMetadata retrieves metadata from a PDF file.
func GetPDFMetadata(inputPDFPath string) error {
	cmd := exec.Command("pdfinfo", inputPDFPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error getting PDF metadata: %v", err)
	}

	return nil
}
