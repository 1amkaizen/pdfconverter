// pdfmetadata.go

package pdfconverter

import (
    "fmt"
    "os/exec"
)

// GetPDFMetadata retrieves metadata from a PDF file.
func GetPDFMetadata(pdfPath string) error {
    cmd := exec.Command("pdfinfo", pdfPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("error getting PDF metadata: %v", err)
    }

    return nil
}
