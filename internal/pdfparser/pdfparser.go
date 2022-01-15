package pdfparser

import (
	"errors"
	"io/ioutil"

	"github.com/davidbyttow/govips/v2/vips"
)

const (
	writePermission = 0644
)

var EmptyPdfBytesError = errors.New("pdf bytes not set")

type PdfParser struct {
	pdfBytes []byte
}

func New() *PdfParser {
	return &PdfParser{}
}

func (s *PdfParser) LoadPdfBytes(pathFileName string) error {
	pdfBytes, err := ioutil.ReadFile(pathFileName)
	if err != nil {
		return err
	}

	s.pdfBytes = pdfBytes
	return nil
}

func (s *PdfParser) SetPdfBytes(pdfBytes []byte) {
	s.pdfBytes = pdfBytes
}

func (s *PdfParser) ConvertToFile(pathFileName string) error {
	imageBytes, err := s.ConvertToBytes()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(pathFileName, imageBytes, writePermission)
}

func (s *PdfParser) ConvertToBytes() ([]byte, error) {
	if len(s.pdfBytes) == 0 {
		return nil, EmptyPdfBytesError
	}

	vips.Startup(nil)
	defer vips.Shutdown()

	params := &vips.ImportParams{}
	params.SvgUnlimited.Set(true)
	image, err := vips.LoadImageFromBuffer(s.pdfBytes, params)
	if err != nil {
		return nil, err
	}

	err = image.Resize(5, vips.KernelMitchell)
	if err != nil {
		return nil, err
	}

	err = image.AutoRotate()
	if err != nil {
		return nil, err
	}

	ep := &vips.PngExportParams{
		Compression: 0,
		Quality:     100,
		Interlace:   true,
		Palette:     true,
		Bitdepth:    2,
	}

	imageBytes, _, err := image.ExportPng(ep)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}
