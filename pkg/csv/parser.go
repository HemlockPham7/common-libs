package csv

import (
	"mime/multipart"

	"github.com/gocarina/gocsv"
)

func ParseFromMultipartFile(src *multipart.FileHeader, data any) error {
	file, err := src.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.Unmarshal(file, data); err != nil {
		return err
	}

	return nil
}
