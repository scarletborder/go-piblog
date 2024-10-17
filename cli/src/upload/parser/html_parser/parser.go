package htmlparser

import (
	"bytes"
	"fmt"
	"piblog/src/model"
	"strings"

	"golang.org/x/net/html"
)

// 返回<title>的内容和<body>内部<div>的所有内容
func extractTitleAndDiv(htmlContent string) (string, string) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return "", ""
	}

	// 获取<title>节点
	titleNode := findNodeByTag(doc, "title")
	var title string
	if titleNode != nil && titleNode.FirstChild != nil {
		title = titleNode.FirstChild.Data // 获取<title>中的文本
	} else {
		title = "Title not found"
	}

	// 获取<div>节点
	divNode := findNodeByTag(doc, "div")
	var divContent string
	if divNode != nil {
		divContent = renderNode(divNode) // 获取<div>的完整HTML
	} else {
		divContent = "Div not found"
	}

	return title, divContent
}

// 查找给定节点中指定标签名的节点
func findNodeByTag(n *html.Node, tagName string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tagName {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findNodeByTag(c, tagName); result != nil {
			return result
		}
	}
	return nil
}

// 将HTML节点转换为字符串
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	err := html.Render(&buf, n)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ParseDocument(input string) model.Document {
	title, content := extractTitleAndDiv(input)
	brief := content[0:Min(len(content), 32)]

	return model.Document{
		Title:   title,
		Content: content,
		Brief:   brief,
	}

}
