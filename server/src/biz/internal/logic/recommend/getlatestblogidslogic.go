package recommend

import (
	"context"

	"biz/internal/logic/archives"
	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestBlogIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLatestBlogIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestBlogIdsLogic {
	return &GetLatestBlogIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLatestBlogIdsLogic) GetLatestBlogIds() (resp *types.BlogIdsResp, err error) {
	latestRecords, err := archives.GetLatestBlogIds(l.ctx, *l.svcCtx, 0)
	if err != nil {
		return
	}
	resp = &types.BlogIdsResp{IDs: latestRecords}
	return
}
