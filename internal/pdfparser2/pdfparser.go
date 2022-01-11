package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	byteSlices, err := ioutil.ReadFile("internal/assets/pdfs/sample1.pdf")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}

	importParams := vips.NewImportParams()
	importParams.JpegShrinkFactor.Set(0)
	importParams.Page.Set(0)
	importParams.NumPages.Set(3)
	importParams.Density.Set(300)

	image, err := vips.LoadImageFromBuffer(byteSlices, importParams)
	checkError(err)
	ep := &vips.PngExportParams{
		Compression: 0,
		Quality:     100,
		Interlace:   true,
		Palette:     true,
		Bitdepth:    2,
	}

	image1bytes, _, err := image.ExportPng(ep)
	err = ioutil.WriteFile("output3.png", image1bytes, 0644)
	checkError(err)
}
