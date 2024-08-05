/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package logic

import (
	"github.com/zeromicro/go-zero/core/conf"
	"IM-CHAT/apps/user/rpc/internal/config"
	"IM-CHAT/apps/user/rpc/internal/svc"
	"path/filepath"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}
