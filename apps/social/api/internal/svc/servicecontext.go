package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"IM-CHAT/apps/social/api/internal/config"
	"IM-CHAT/apps/social/rpc/socialclient"
	"IM-CHAT/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	socialclient.Social
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
