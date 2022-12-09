package user

import (
	"net/http"

	"zeros/app/usercenter/cmd/api/internal/logic/user"
	"zeros/app/usercenter/cmd/api/internal/svc"
	"zeros/app/usercenter/cmd/api/internal/types"
	"zeros/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
