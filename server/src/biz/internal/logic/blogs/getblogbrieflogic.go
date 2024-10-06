package blogs

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlogBriefLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// type 

func NewGetBlogBriefLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogBriefLogic {
	return &GetBlogBriefLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取单个blog

func (l *GetBlogBriefLogic) GetBlogBrief(req *types.BlogIdReq) (resp *types.BlogBriefResp, err error) {
	res, err := l.svcCtx.BlogModel.FindOne(l.ctx, req.ID)
	if err != nil {
		// 研究出go-zero传入ctx中是否包含有更多信息再做处理
		// 其实也不用,因为自带了,log with context会解决这个问题
		return
	}
	resp = &types.BlogBriefResp{
		ID:    req.ID,
		Brief: res.Brief,
		Tags:  res.Tags,
	}
	return
}
