// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	blogs "biz/internal/handler/blogs"
	test "biz/internal/handler/test"
	"biz/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/content/:id",
				Handler: blogs.GetBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: blogs.GetBlogBriefsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info/:id",
				Handler: blogs.GetBlogBriefHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/blog"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: test.PingHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/version",
				Handler: test.VersionHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
