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

// ConvertTextToPDF converts a text file to PDF.
func ConvertTextToPDF(inputTextPath, outputPDFPath string) error {
    cmd := exec.Command("pdftotext", inputTextPath, outputPDFPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("error converting text to PDF: %v", err)
    }

    return nil
}
