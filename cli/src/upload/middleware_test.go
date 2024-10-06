package upload_test

import (
	"piblog/src/upload"
	"testing"
)

func TestGlobalComponent(t *testing.T) {
	input := `
# Example Title
[tag ex1, tag ex2]
The text following tags will be recognized as brief text. Brief text ends when meet newline

Some texts which only display in main text

## Example subtitle 1

Some texts which only display in main text

## Example subtitle 2

### Some text

Some texts which only display in main text
`
	doc := upload.ParseDocument(input)
	myUploader := upload.GlobalUpload.Copy()
	myUploader.Use(upload.DocumentHandler(upload.LogHandler).ToMiddleWare())
	myUploader.Process(doc)
}
