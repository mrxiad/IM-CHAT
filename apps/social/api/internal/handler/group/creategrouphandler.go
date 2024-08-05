package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"IM-CHAT/apps/social/api/internal/logic/group"
	"IM-CHAT/apps/social/api/internal/svc"
	"IM-CHAT/apps/social/api/internal/types"
)

func CreateGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := group.NewCreateGroupLogic(r.Context(), svcCtx)
		resp, err := l.CreateGroup(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
