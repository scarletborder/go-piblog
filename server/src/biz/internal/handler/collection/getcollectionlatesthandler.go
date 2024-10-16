package collection

import (
	"net/http"

	"biz/internal/logic/collection"
	"biz/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCollectionLatestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := collection.NewGetCollectionLatestLogic(r.Context(), svcCtx)
		resp, err := l.GetCollectionLatest()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
