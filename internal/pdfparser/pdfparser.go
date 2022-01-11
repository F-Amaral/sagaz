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

	params := &vips.ImportParams{}
	params.SvgUnlimited.Set(true)
	image1, err := vips.LoadImageFromBuffer(byteSlices, params)
	checkError(err)

	err = image1.Resize(5, vips.KernelMitchell)
	checkError(err)

	// Rotate the picture upright and reset EXIF orientation tag
	err = image1.AutoRotate()
	checkError(err)

	ep := &vips.PngExportParams{
		Compression: 0,
		Quality:     100,
		Interlace:   true,
		Palette:     true,
		Bitdepth:    2,
	}

	image1bytes, _, err := image1.ExportPng(ep)
	err = ioutil.WriteFile("output.png", image1bytes, 0644)
	checkError(err)
}
