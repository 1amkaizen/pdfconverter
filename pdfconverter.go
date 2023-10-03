// pdfconverter.go

package pdfconverter

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintConvertedText() error {
	cmd := exec.Command("cat", "output.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error printing converted text: %v", err)
	}

	return nil
}

func ConvertPDFToText(inputPDFPath string) error {
	cmd := exec.Command("pdftotext", inputPDFPath, "output.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting PDF to text: %v", err)
	}

	return nil
}
