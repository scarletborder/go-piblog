package collection

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionBriefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCollectionBriefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionBriefLogic {
	return &GetCollectionBriefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCollectionBriefLogic) GetCollectionBrief(req *types.CollectionBriefReq) (resp *types.CollectionBriefResp, err error) {
	// todo: add your logic here and delete this line

	return
}
