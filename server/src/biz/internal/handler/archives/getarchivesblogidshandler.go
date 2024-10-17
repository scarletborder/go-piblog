package archives

import (
	"net/http"

	"biz/internal/logic/archives"
	"biz/internal/svc"
	"biz/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArchivesBlogIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArchivesBlogIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := archives.NewGetArchivesBlogIdsLogic(r.Context(), svcCtx)
		resp, err := l.GetArchivesBlogIds(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Header().Set("Cache-Control", "max-age=60, public")
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
