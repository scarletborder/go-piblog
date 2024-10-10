package logic

import (
	"context"
	"errors"

	"host/internal/svc"
	"host/pb/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlogLogic {
	return &CreateBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBlogLogic) CreateBlog(in *host.UploadBlogReq) (*host.UploadBlogResp, error) {
	// todo: add your logic here and delete this line

	return &host.UploadBlogResp{Status: false, Msg: "This method is not available"}, errors.New("This method is not available")
}
