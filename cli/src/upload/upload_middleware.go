package upload

import (
	"context"
	"fmt"
	"piblog/src/config"
	"piblog/src/hostclient"
	"piblog/src/model"
	"piblog/src/pb/host"

	"github.com/zeromicro/go-zero/zrpc"
)

// Register all used middleware here
var GlobalUpload = model.NewDefaultUploader()

func init() {
	GlobalUpload.Use(model.DocumentHandler(model.LogHandler).ToMiddleWare())

	// TODO: add other middleware here
	// ...
	// GlobalUpload.Use(LogDetailMiddleWare)
}

func LogDetailMiddleWare(next model.DocumentHandler) model.DocumentHandler {
	return model.DocumentHandler(func(d model.Document) {
		fmt.Printf("detail information: %s", d.ToDebugString())
		next(d)
	})
}

func DBUpdateHandler(d model.Document) {
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

// TODO: collection upload
