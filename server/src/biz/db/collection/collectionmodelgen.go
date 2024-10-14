// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type collectionModel interface {
	Insert(ctx context.Context, data *Collection) error
	FindOne(ctx context.Context, id string) (*Collection, error)
	Update(ctx context.Context, data *Collection) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultCollectionModel struct {
	conn *mon.Model
}

func newDefaultCollectionModel(conn *mon.Model) *defaultCollectionModel {
	return &defaultCollectionModel{conn: conn}
}

func (m *defaultCollectionModel) Insert(ctx context.Context, data *Collection) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultCollectionModel) FindOne(ctx context.Context, id string) (*Collection, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Collection

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCollectionModel) Update(ctx context.Context, data *Collection) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultCollectionModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
