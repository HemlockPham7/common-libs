package csv

import (
	"bytes"
	"io"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CreateTestMultipartRequest creates a multipart request for mock testing containing a test.csv file.
// The provided content is written to the test.csv file.
func CreateTestMultipartRequest(t *testing.T, content string) (*multipart.Writer, *bytes.Buffer) {
	body := &bytes.Buffer{}             // Create a temporary memory buffer in Ram to store the request body
	writer := multipart.NewWriter(body) // Create a new multipart writer to write data into the body

	// Create a temporary file to store content in multipart request
	part, err := writer.CreateFormFile("file", "test.csv") // part: place to store the content CSV file into
	assert.NoError(t, err)

	_, err = io.Copy(part, bytes.NewBuffer([]byte(content)))
	assert.NoError(t, err)

	err = writer.Close()
	assert.NoError(t, err)

	return writer, body
}
