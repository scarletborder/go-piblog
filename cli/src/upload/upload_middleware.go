package upload

import (
	"context"
	"fmt"
	"piblog/src/config"
	"piblog/src/hostclient"
	"piblog/src/pb/host"
	"strings"

	"github.com/zeromicro/go-zero/zrpc"
)

// Register all used middleware here
func init() {
	GlobalUpload.documentHandler = LogHandler

	// TODO: add other middleware here
	// ...
	// GlobalUpload.Use(LogDetailMiddleWare)
}

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

var GlobalUpload UploadComponent

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

// Default document handler
func LogHandler(d Document) {
	fmt.Printf("Uploaded document titled %s\nWith tags: %s\n",
		d.Title, strings.Join(d.Tags, "||"))

	myMap, ok := d.context.Value(DocumentContextMap{}).(map[string]interface{})
	if ok && len(myMap) > 0 {
		fmt.Printf("With context:\n")
		for key, value := range myMap {
			fmt.Printf("%s: %v\n", key, value)
		}
	}

}

func LogDetailMiddleWare(next DocumentHandler) DocumentHandler {
	return DocumentHandler(func(d Document) {
		fmt.Printf("detail information: %s", d.ToDebugString())
		next(d)
	})
}

func DBUpdateHandler(d Document) {
	// TODO: 重新使用默认
	tmp := config.GlobalConfig.HostRpc
	tmp.Timeout = 4000000
	cli := zrpc.MustNewClient(tmp)

	hostClient := hostclient.NewHost(cli)

	in := &host.UploadBlogReq{
		Title:   d.Title,
		Tags:    d.Tags,
		Brief:   d.Brief,
		Content: d.Content,
	}
	resp, err := hostClient.UpdateBlog(context.Background(), in)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Success Update, ID:%s", resp.Id)
	}
}
