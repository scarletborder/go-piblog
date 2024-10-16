package archives

import (
	"context"
	"errors"

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
	page_num, err := GetArchivePageNumber(l.ctx, l.svcCtx.BlogModel.GetConn(), l.svcCtx.Rds)
	if err != nil {
		return nil, err
	}
	if req.Page >= int(page_num) {
		return nil, errors.New("your request exceed max page number")
	}
	res, err := GetArchivesBlogIds(l.ctx, *l.svcCtx, int64(req.Page))
	if err != nil {
		return nil, err
	}

	return &types.ArchivesBlogIdsResp{
		MaxPages: int(page_num),
		IDs:      res,
	}, nil
}
