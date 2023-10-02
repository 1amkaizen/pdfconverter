# PDF Converter Library

Library untuk mengkonversi file PDF menjadi teks menggunakan perintah pdftotext.

## Penggunaan

1. Install library ini dengan menambahkannya sebagai dependency pada file `go.mod` proyek Anda:
```
go get github.com/1amkaizen/pdfconverter
```

2. Import library ini di proyek Anda:
```go
import "github.com/1amkaizen/pdfconverter"
```
3. Panggil fungsi ConvertPDFToText untuk mengkonversi file PDF menjadi teks:
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
```
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

### Kontribusi
Jika Anda ingin berkontribusi pada proyek ini, silakan buka CONTRIBUTING.md untuk panduan kontribusi.
