// pdfconverter.go

package pdfconverter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// ConvertPDFToText converts a PDF file to text and returns the text.
func ConvertPDFToText(inputPDFPath string) (string, error) {
	cmd := exec.Command("pdftotext", inputPDFPath, "-")
	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &outBuffer

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error converting PDF to text: %v", err)
	}

	return outBuffer.String(), nil
}
