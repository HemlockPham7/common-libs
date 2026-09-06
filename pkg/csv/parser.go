package csv

import (
	"mime/multipart"

	"github.com/gocarina/gocsv"
)

// ParseFromMultipartFile parses CSV data from a multipart file into the provided destination.
//
// Parameters:
//   - src: the uploaded multipart file containing CSV data.
//   - data: the destination where the parsed CSV records are stored.
//
// Returns:
//   - An error if the file cannot be opened or the CSV data cannot be parsed.
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
