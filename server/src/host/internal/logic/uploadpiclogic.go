package logic

import (
	"context"

	"host/internal/svc"
	"host/pb/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPicLogic {
	return &UploadPicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPicLogic) UploadPic(in *host.UploadPicReq) (*host.UploadPicResp, error) {
	// todo: add your logic here and delete this line

	return &host.UploadPicResp{}, nil
}
