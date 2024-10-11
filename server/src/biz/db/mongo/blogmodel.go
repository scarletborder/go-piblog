package model

import (
	"biz/internal/config"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ BlogModel = (*customBlogModel)(nil)

type (
	// BlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogModel.
	BlogModel interface {
		blogModel
		GetConn() *monc.Model
	}

	customBlogModel struct {
		*defaultBlogModel
	}
)

// NewBlogModel returns a model for the mongo.
func NewBlogModel(url, db, collection string, c cache.CacheConf) BlogModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customBlogModel{
		defaultBlogModel: newDefaultBlogModel(conn),
	}
}

func NewBlogModelByCfg(url string, cfg config.MongoConfig, c cache.CacheConf) BlogModel {
	return NewBlogModel(url, cfg.BlogModel.DbName, cfg.BlogModel.CollectionName, c)
}

func (m *defaultBlogModel) GetConn() *monc.Model {
	return m.conn
}
