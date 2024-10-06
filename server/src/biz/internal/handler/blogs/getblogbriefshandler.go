package blogs

import (
	"net/http"

	"biz/internal/logic/blogs"
	"biz/internal/svc"
	"biz/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBlogBriefsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BlogIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blogs.NewGetBlogBriefsLogic(r.Context(), svcCtx)
		resp, err := l.GetBlogBriefs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
