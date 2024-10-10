package logic

import (
	"context"
	"errors"

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
	_id := in.GetId()
	if _id == "" {
		return &host.DeleteBlogResp{Status: false, Msg: "no id of blog is provided"}, errors.New("no id of blog is provided")
	}
	del_num, err := l.svcCtx.BlogModel.Delete(l.ctx, _id)
	if err != nil {
		return &host.DeleteBlogResp{Status: false, Msg: err.Error()}, err
	}

	if del_num < 1 {
		return &host.DeleteBlogResp{Status: true, Msg: "no suitable blog with the id is found"}, nil
	}

	return &host.DeleteBlogResp{Status: true, Msg: "success"}, nil
}
