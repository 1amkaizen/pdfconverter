# PDF Converter Library

Library untuk mengkonversi file PDF menjadi teks menggunakan perintah pdftotext.

## Penggunaan

1. Install poppler-utils

```
sudo apt-get install poppler-utils
```
 

3. Install library ini dengan menambahkannya sebagai dependency pada file `go.mod` proyek Anda:
```
go get github.com/1amkaizen/pdfconverter
```

3. Import library ini di proyek Anda:
```go
import "github.com/1amkaizen/pdfconverter"
```
4. Panggil fungsi ConvertPDFToText untuk mengkonversi file PDF menjadi teks:
```go
pdfFilePath := "path/to/input.pdf"
outputFilePath := "path/to/output.txt"

err := pdfconverter.ConvertPDFToText(pdfFilePath, outputFilePath)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Println("PDF berhasil diubah menjadi teks.")

```

### Contoh Penggunaan
Berikut adalah contoh penggunaan library untuk mengkonversi file PDF menjadi teks:
```go
package main

import (
	"fmt"
	"github.com/1amkaizen/pdfconverter"
)

func main() {
	pdfFilePath := "input.pdf"
	outputFilePath := "output.txt"

	err := pdfconverter.ConvertPDFToText(pdfFilePath, outputFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Konversi PDF selesai.")
}

```

### Contoh lain
```go
package main

import (
	"fmt"
	"os"
	"github.com/1amkaizen/pdfconverter" // Sesuaikan dengan path package yang sesuai
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: main inputPDFPath outputTextPath")
		return
	}

	inputPDFPath := os.Args[1]
	outputTextPath := os.Args[2]

	err := pdfconverter.ConvertPDFToText(inputPDFPath, outputTextPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("PDF converted to text successfully.")
}
```
Untuk menjalankannya:
```
go run main.go input.pdf output.txt

```

### Kontribusi
Jika Anda ingin berkontribusi pada proyek ini, silakan buka [CONTRIBUTING.md](https://github.com/1amkaizen/pdfconverter/edit/main/CONTRIBUTING.md ) untuk panduan kontribusi.
