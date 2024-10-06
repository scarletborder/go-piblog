package config

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MongoConfig MongoConfig

	AuthApps []AuthApp
}

type AuthApp struct {
	App   string
	Token string // 默认的redis中对rpc服务的access token
}

type MongoConfig struct {
	Uri, Host string
	Port      int
	BlogModel MongoModel
	NodeConfs []cache.NodeConf
}

type MongoModel struct {
	DbName, CollectionName string
}

func (m MongoConfig) ToUri() (uri string) {
	uri = fmt.Sprintf("%s@%s:%d", m.Uri, m.Host, m.Port)
	return
}

// TODO: delete this
// 只是想试试RpcClientConf里能不能设置auth
// 可以
func init() {
	// zrpc.MustNewClient()
	// zrpc.NewClient(zrpc.RpcClientConf{Endpoints: []string{127.0.0.1:8080}})
}
