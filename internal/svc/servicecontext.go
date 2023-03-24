package svc

import (
	"github.com/jarcn/shoutapi/internal/config"

	"github.com/jarcn/shoutrpc/transformer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer //配置远程服务
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), //配置远程服务
	}
}
