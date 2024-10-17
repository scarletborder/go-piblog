package markdownparser

import (
	"piblog/src/model"
	"strings"
)


// 对文件内容进行解析
func  ParseDocument(input string) model.Document {
	lines := strings.Split(input, "\n")
	var doc model.Document

	// 解析 Title
	for i, line := range lines {
		line = strings.TrimPrefix(line, "\ufeff")
		if strings.HasPrefix(line, "# ") {
			doc.Title = strings.TrimSpace(strings.TrimPrefix(line, "# "))
			lines = lines[i+1:]
			break
		}
	}

	// 解析 Tags
	lines[0] = strings.TrimSpace(lines[0])
	if len(lines) > 0 && strings.HasPrefix(lines[0], "[") && strings.HasSuffix(lines[0], "]") {
		tagLine := strings.TrimSpace(lines[0])
		tagLine = strings.Trim(tagLine, "[]")
		pre_tags := strings.Split(tagLine, ",")
		for _, tag_item := range pre_tags {
			doc.Tags = append(doc.Tags, strings.TrimSpace(tag_item))
		}
		lines = lines[1:]
	} else {
		doc.Tags = []string{"blank tag"}
	}

	// 解析 Brief
	var briefLines []string
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "---" { // 遇到---停止解析 Brief
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
