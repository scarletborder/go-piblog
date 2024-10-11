package blogs

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogLogic {
	return &GetBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlogLogic) GetBlog(req *types.BlogIdReq) (resp *types.BlogDetailResp, err error) {
	res, err := l.svcCtx.BlogModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return
	}
	resp = &types.BlogDetailResp{
		ID:      req.ID,
		Title:   res.Title,
		Content: res.Content,
		Tags:    res.Tags,
	}

	return
}
