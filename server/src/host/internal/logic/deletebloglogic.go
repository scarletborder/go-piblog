package logic

import (
	"context"

	"host/internal/svc"
	"host/pb/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlogLogic {
	return &DeleteBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBlogLogic) DeleteBlog(in *host.DeleteBlogReq) (*host.DeleteBlogResp, error) {
	// todo: add your logic here and delete this line

	return &host.DeleteBlogResp{}, nil
}
