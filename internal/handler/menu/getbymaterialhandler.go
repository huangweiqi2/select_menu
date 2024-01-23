package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"select_menu/internal/logic/menu"
	"select_menu/internal/svc"
	"select_menu/internal/types"
)

func GetByMaterialHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetByMaterialRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetByMaterialLogic(r.Context(), svcCtx)
		resp, err := l.GetByMaterial(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
