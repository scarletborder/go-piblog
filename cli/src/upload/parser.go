package upload

import (
	"context"
	"fmt"
	"strings"
)

type Document struct {
	Title   string
	Tags    []string
	Brief   string
	Content string

	context context.Context
}

func (d Document) ToDebugString() string {
	return fmt.Sprintf("Title: %s\nTags: %v\nBrief: %s\nContent: \n%s\n",
		d.Title, d.Tags, d.Brief, d.Content)
}

// 在上传前对Blog进行处理（例如将tags进行统计，创建评论条目）
type DocumentHandler func(Document)

// Context中需要打印的内容
type DocumentContextMap struct{}

// 对文件内容进行解析
func ParseDocument(input string) Document {
	lines := strings.Split(input, "\n")
	var doc Document
	doc.context = context.WithValue(context.Background(),
		DocumentContextMap{}, map[string]interface{}{})

	// 解析 Title
	for i, line := range lines {
		if strings.HasPrefix(line, "# ") {
			doc.Title = strings.TrimSpace(strings.TrimPrefix(line, "# "))
			lines = lines[i+1:]
			break
		}
	}

	// 解析 Tags
	if len(lines) > 0 && strings.HasPrefix(lines[0], "[") && strings.HasSuffix(lines[0], "]") {
		tagLine := strings.TrimSpace(lines[0])
		tagLine = strings.Trim(tagLine, "[]")
		doc.Tags = strings.Split(tagLine, ", ")
		lines = lines[1:]
	}

	// 解析 Brief
	var briefLines []string
	for i, line := range lines {
		if line == "" { // 遇到空行停止解析 Brief
			lines = lines[i+1:]
			break
		}
		briefLines = append(briefLines, strings.TrimSpace(line))
	}
	doc.Brief = strings.Join(briefLines, " ")

	// 剩余部分全部作为 Content
	doc.Content = strings.Join(lines, "\n")

	return doc
}
