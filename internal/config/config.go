package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Transform zrpc.RpcClientConf //添加远程服务依赖
	Add   zrpc.RpcClientConf     // 手动代码
    Check zrpc.RpcClientConf     // 手动代码
}
