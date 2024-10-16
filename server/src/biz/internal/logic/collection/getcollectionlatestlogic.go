package collection

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionLatestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCollectionLatestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionLatestLogic {
	return &GetCollectionLatestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCollectionLatestLogic) GetCollectionLatest() (resp *types.CollectionLatestResp, err error) {
	// todo: add your logic here and delete this line

	return
}
