package menu

import (
	"context"
	"github.com/samber/lo"
	"select_menu/helper"
	"select_menu/models"

	"select_menu/internal/svc"
	"select_menu/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RandomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRandomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RandomLogic {
	return &RandomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RandomLogic) Random(req *types.RandomRequest) (resp *types.RandomResponse, err error) {
	var foods []models.Food
	err = models.DB.Model(new(models.Food)).Find(&foods).Error
	if err != nil {
		return
	}
	hots := make([]types.FoodResponse, 0, len(foods))
	colds := make([]types.FoodResponse, 0, len(foods))
	soups := make([]types.FoodResponse, 0, len(foods))

	for _, food := range foods {
		if food.Status.IsHot() {
			hots = append(hots, food.Response())
		} else if food.Status.IsCold() {
			colds = append(colds, food.Response())
		} else {
			soups = append(soups, food.Response())
		}
	}

	if req.HotNum > 0 {
		if len(hots) < req.HotNum {
			resp.Foods = append(resp.Foods, hots...)
		} else {
			resp.Foods = append(resp.Foods, helper.SliceRandomN(hots, req.HotNum)...)
			//	log.Printf()
		}
	}
	if req.ColdNum > 0 {
		if len(colds) < req.ColdNum {
			resp.Foods = append(resp.Foods, colds...)
		} else {
			resp.Foods = append(resp.Foods, helper.SliceRandomN(colds, req.ColdNum)...)
		}
	}
	if req.SoupNum > 0 {
		if len(soups) < req.SoupNum {
			resp.Foods = append(resp.Foods, soups...)
		} else {
			resp.Foods = append(resp.Foods, helper.SliceRandomN(soups, req.SoupNum)...)
		}
	}

	//materials := make([]string, 0)
	//for _, food := range resp.Foods {
	//	materials = append(materials, food.Material...)
	//}
	//去重

	f := func(item types.FoodResponse, index int) []string {
		return item.Material
	}
	flatMap := lo.FlatMap(resp.Foods, f)

	resp.Materials = lo.Uniq(flatMap)
	//m := make(map[string]struct{}, len(materials))
	//for _, material := range materials {
	//	if _, ok := m[material]; ok {
	//		continue
	//	}
	//
	//	m[material] = struct{}{}
	//	resp.Materials = append(resp.Materials, material)
	//
	//}

	return
}
