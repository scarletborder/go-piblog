package collection

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCollectionPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionPageLogic {
	return &GetCollectionPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCollectionPageLogic) GetCollectionPage(req *types.CollectionPageReq) (resp *types.CollectionPageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
