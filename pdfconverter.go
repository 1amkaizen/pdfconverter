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

	// Call the new feature
	features.NewFeature()

	return nil
}

