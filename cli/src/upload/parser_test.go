package upload_test

// USAGE:
// go test -v : show detailed output

import (
	"piblog/src/upload"
	"testing"
)

func TestParseDocument(t *testing.T) {
	// 示例文本
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
	t.Logf("struct output: %v", doc.ToDebugString())
}
