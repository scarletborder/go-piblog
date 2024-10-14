package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ CollectionModel = (*customCollectionModel)(nil)

type (
	// CollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectionModel.
	CollectionModel interface {
		collectionModel
	}

	customCollectionModel struct {
		*defaultCollectionModel
	}
)

// NewCollectionModel returns a model for the mongo.
func NewCollectionModel(url, db, collection string) CollectionModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customCollectionModel{
		defaultCollectionModel: newDefaultCollectionModel(conn),
	}
}
