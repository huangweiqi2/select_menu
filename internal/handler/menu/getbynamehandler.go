package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"select_menu/internal/logic/menu"
	"select_menu/internal/svc"
	"select_menu/internal/types"
)

func GetByNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetByNameRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetByNameLogic(r.Context(), svcCtx)
		resp, err := l.GetByName(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
