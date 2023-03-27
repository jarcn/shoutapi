package svc

import (
	"github.com/jarcn/shoutapi/internal/config"

	"github.com/jarcn/shoutrpc/transformer"
	"github.com/jarcn/bookadd/adder"
	"github.com/jarcn/bookcheck/checker"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer //配置远程服务
	Adder   adder.Adder          // 手动代码
    Checker checker.Checker      // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), //配置远程服务
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),         // 手动代码
        Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),   // 手动代码
	}
}
