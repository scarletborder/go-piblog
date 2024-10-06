package logic

import (
	"context"

	"host/internal/svc"
	"host/pb/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePicLogic {
	return &DeletePicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePicLogic) DeletePic(in *host.DeleteBlogReq) (*host.DeletePicResp, error) {
	// todo: add your logic here and delete this line

	return &host.DeletePicResp{}, nil
}
