package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	PRpc1    zrpc.RpcClientConf
	EsConfig EsConfig
}

// ES配置
type EsConfig struct {
	Addresses  []string
	ConfigPath string
	Indexes    []string
	Username   string
	Password   string
}
