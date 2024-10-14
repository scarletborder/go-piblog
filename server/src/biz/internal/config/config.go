package config

// declare all Config structure here, and fill them in .yaml file

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

// 定义服务端的更多与服务无关的细节信息
type Detail struct {
	Version int64
}

type MongoConfig struct {
	Uri, Host string
	Port      int
	BlogModel MongoModel
	NodeConfs []cache.NodeConf
}

// 各种数据库中模型的来源
// 博客文章模型
type MongoModel struct {
	DbName, CollectionName string
}

func (m MongoConfig) ToUri() (uri string) {
	uri = fmt.Sprintf("%s@%s:%d", m.Uri, m.Host, m.Port)
	return
}

// 用户请求API的细节
type APILimit struct {
	MaxBlogNumber int
}

type Config struct {
	rest.RestConf
	APILimit    APILimit
	Detail      Detail
	MongoConfig MongoConfig
}
