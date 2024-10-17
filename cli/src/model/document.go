package model

import (
	"context"
	"fmt"
)

type Document struct {
	Title   string
	Tags    []string
	Brief   string
	Content string

	Context context.Context
}

func NewDocument() Document {
	return Document{
		Context: context.WithValue(context.Background(),
			DocumentContextMap{}, map[string]interface{}{}),
	}
}

func (d Document) ToDebugString() string {
	return fmt.Sprintf("Title: %s\nTags: %v\nBrief: %s\nContent: \n%s\n",
		d.Title, d.Tags, d.Brief, d.Content)
}

// Context中需要打印的内容 Key
type DocumentContextMap struct{}

func (d Document)AddLogThing(k string, v interface{}) {
	myMap, _ := d.Context.Value(DocumentContextMap{}).(map[string]interface{})
	myMap[k] = v
}