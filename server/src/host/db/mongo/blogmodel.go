package model

import (
	"context"
	"host/internal/config"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
)

var _ BlogModel = (*customBlogModel)(nil)

type (
	// BlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogModel.
	BlogModel interface {
		blogModel
		UpdateByTitle(ctx context.Context, data *Blog) (*Blog, error)
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

func (m *defaultBlogModel) UpdateByTitle(ctx context.Context, data *Blog) (*Blog, error) {
	var res Blog
	filter := bson.M{
		"Title": bson.M{"$regex": data.Title}, // 原样匹配
	}
	err := m.conn.Model.FindOne(ctx, &res, filter)
	if monc.ErrNotFound != err {
		// 原已有
		data.ID = res.ID
		_, err = m.Update(ctx, data)
		return data, err
	}
	// Insert
	err = m.Insert(ctx, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
