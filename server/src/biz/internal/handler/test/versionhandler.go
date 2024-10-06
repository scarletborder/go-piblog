package test

import (
	"net/http"

	"biz/internal/logic/test"
	"biz/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := test.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
