package collection

import (
	"net/http"

	"biz/internal/logic/collection"
	"biz/internal/svc"
	"biz/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCollectionBriefHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectionBriefReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := collection.NewGetCollectionBriefLogic(r.Context(), svcCtx)
		resp, err := l.GetCollectionBrief(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
