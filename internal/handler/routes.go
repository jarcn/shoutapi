// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/jarcn/shoutapi/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/shorten",
				Handler: ShortenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/expand",
				Handler: ExpandHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/add",
				Handler: AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/check",
				Handler: CheckHandler(serverCtx),
			},
		},
	)
}
