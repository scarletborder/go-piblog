package model

import (
	"fmt"
	"strings"
)

// 在上传前对Blog进行处理（例如将tags进行统计，创建评论条目）
type DocumentHandler func(Document)

// 根据若干规则处理即将上传的博客文件
type UploadMiddleWare func(DocumentHandler) DocumentHandler

type UploadComponent struct {
	documentHandler DocumentHandler
}

// 新建UploadComponent, Handler指向原来的Handler
func (u UploadComponent) Copy() UploadComponent {
	return UploadComponent{
		documentHandler: u.documentHandler,
	}
}

func (u *UploadComponent) Use(middle_ware UploadMiddleWare) {
	u.documentHandler = middle_ware(u.documentHandler)
}

func (u *UploadComponent) Process(d Document) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("博客Document处理捕获到 panic: %v", r)
		}
	}()

	if u.documentHandler != nil {
		u.documentHandler(d)
	}

	return
}

func (hand DocumentHandler) ToMiddleWare() UploadMiddleWare {
	return UploadMiddleWare(func(next DocumentHandler) DocumentHandler {
		return DocumentHandler(func(d Document) {
			hand(d)
			next(d)
		})
	})
}

func nilHandler(d Document) {
	// nil here
}

func LogHandler(d Document) {
	fmt.Printf("Uploaded document titled %s\nWith tags: %s\n",
		d.Title, strings.Join(d.Tags, "||"))

	myMap, ok := d.Context.Value(DocumentContextMap{}).(map[string]interface{})
	if ok && len(myMap) > 0 {
		fmt.Printf("With context:\n")
		for key, value := range myMap {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

func NewDefaultUploader() UploadComponent {
	var DefaultUpload UploadComponent
	DefaultUpload.documentHandler = nilHandler
	return DefaultUpload
}
