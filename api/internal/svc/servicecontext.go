package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/zrpc"
)

// 通过ServiceContext在不同业务逻辑之间传递依赖
type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder     // 1.rpc client call entry
	Checker checker.Checker // 2.rpc client call entry
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),       // 1.
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)), // 2.
	}
}
