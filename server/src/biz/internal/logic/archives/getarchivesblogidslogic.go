package archives

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArchivesBlogIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArchivesBlogIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArchivesBlogIdsLogic {
	return &GetArchivesBlogIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArchivesBlogIdsLogic) GetArchivesBlogIds(req *types.ArchivesBlogIdsReq) (resp *types.ArchivesBlogIdsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
