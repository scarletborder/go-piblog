package blogs

import (
	"context"

	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetBlogBriefsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlogBriefsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogBriefsLogic {
	return &GetBlogBriefsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
获取多个id博文的brief
*/
func (l *GetBlogBriefsLogic) GetBlogBriefs(req *types.BlogIdsReq) (resp []types.BlogBriefResp, err error) {
	resp, err = mr.MapReduce(func(source chan<- string) {
		for _, c_id := range req.IDs {
			source <- c_id
		}
	}, func(c_id string, writer mr.Writer[types.BlogBriefResp], cancel func(error)) {
		c_res, c_err := l.svcCtx.BlogModel.FindOne(l.ctx, c_id)
		if c_err != nil {
			cancel(c_err) // 终止所有的逻辑
			return        // 退出当前的逻辑
		}
		c_brief := types.BlogBriefResp{
			ID:    c_id,
			Title: c_res.Title,
			Tags:  c_res.Tags,
			Brief: c_res.Brief,
		}
		writer.Write(c_brief)

	}, func(pipe <-chan types.BlogBriefResp, writer mr.Writer[[]types.BlogBriefResp], cancel func(error)) {
		var resp []types.BlogBriefResp
		for item := range pipe {
			resp = append(resp, item)
		}
		writer.Write(resp)
	})
	return
}
