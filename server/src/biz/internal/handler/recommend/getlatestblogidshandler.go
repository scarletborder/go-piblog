package recommend

import (
	"net/http"

	"biz/internal/logic/recommend"
	"biz/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetLatestBlogIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := recommend.NewGetLatestBlogIdsLogic(r.Context(), svcCtx)
		resp, err := l.GetLatestBlogIds()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
