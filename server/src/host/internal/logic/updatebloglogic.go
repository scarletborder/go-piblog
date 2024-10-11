package logic

import (
	"context"
	model "host/db/mongo"
	"host/internal/svc"
	"host/pb/host"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBlogLogic {
	return &UpdateBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBlogLogic) UpdateBlog(in *host.UploadBlogReq) (*host.UploadBlogResp, error) {
	// todo: add your logic here and delete this line
	blog := &model.Blog{
		Title:   in.GetTitle(),
		Tags:    in.GetTags(),
		Brief:   in.GetBrief(),
		Content: in.GetContent(),
	}
	data, err := l.svcCtx.BlogModel.UpdateByTitle(l.ctx, blog)

	var resp host.UploadBlogResp
	if err != nil {
		resp.Status = false
		resp.Msg = err.Error()
		resp.Id = ""
	} else {
		resp.Status = true
		resp.Msg = "success"
		resp.Id = data.ID.Hex()
	}
	return &resp, err
}
