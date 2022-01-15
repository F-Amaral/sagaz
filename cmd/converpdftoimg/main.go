package main

import (
	"fmt"
	"github.com/johnfercher/sagaz/internal/pdfparser"
)

func main() {
	pdfParser := pdfparser.New()

	err := pdfParser.LoadPdfBytes("internal/assets/pdfs/sample1.pdf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = pdfParser.ConvertToFile("output.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
