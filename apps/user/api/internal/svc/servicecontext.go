package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"IM-CHAT/apps/user/api/internal/config"
	"IM-CHAT/apps/user/rpc/userclient"

	// N * client =》 别名
)

type ServiceContext struct {
	Config config.Config

	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
