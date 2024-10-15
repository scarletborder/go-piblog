package collection

import (
	"net/http"

	"biz/internal/logic/collection"
	"biz/internal/svc"
	"biz/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCollectionPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectionPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := collection.NewGetCollectionPageLogic(r.Context(), svcCtx)
		resp, err := l.GetCollectionPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
