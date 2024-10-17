package blogs

import (
	"net/http"

	"biz/internal/logic/blogs"
	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBlogBriefHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BlogIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blogs.NewGetBlogBriefLogic(r.Context(), svcCtx)
		resp, err := l.GetBlogBrief(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			w.Header().Set("Cache-Control", "max-age=1800, public")
			// TODO: 根据update time来设置			w.Header().Set("Last-Modified", resp)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
