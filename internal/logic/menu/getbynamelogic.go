package menu

import (
	"context"
	"select_menu/internal/errs"
	"select_menu/internal/svc"
	"select_menu/internal/types"
	"select_menu/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByNameLogic {
	return &GetByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByNameLogic) GetByName(req *types.GetByNameRequest) (resp *types.RandomResponse, err error) {
	var foods = make([]models.Food, 0)
	err = models.DB.Model(new(models.Food)).Where("name like ?", "%"+req.Name+"%").Find(&foods).Error
	if err != nil {
		err = errs.QueryModelErr
	}
	resp = new(types.RandomResponse)
	for _, food := range foods {
		resp.Foods = append(resp.Foods, food.Response())
	}
	return
}
