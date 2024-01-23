package menu

import (
	"context"
	"github.com/samber/lo"
	"select_menu/helper"
	"select_menu/internal/errs"
	"select_menu/internal/svc"
	"select_menu/internal/types"
	"select_menu/models"
	"strings"

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
		err = errs.QueryModelErr
		return
	}
	hots := make([]models.Food, 0, len(foods))
	colds := make([]models.Food, 0, len(foods))
	soups := make([]models.Food, 0, len(foods))

	for _, food := range foods {
		if food.Status.IsHot() {
			hots = append(hots, food)
		} else if food.Status.IsCold() {
			colds = append(colds, food)
		} else {
			soups = append(soups, food)
		}
	}

	var resultFoods []models.Food
	if req.HotNum > 0 {
		if len(hots) < req.HotNum {
			resultFoods = append(resultFoods, hots...)
		} else {
			resultFoods = append(resultFoods, helper.SliceRandomN(hots, req.HotNum)...)
		}
	}
	if req.ColdNum > 0 {
		if len(colds) < req.ColdNum {
			resultFoods = append(resultFoods, colds...)
		} else {
			resultFoods = append(resultFoods, helper.SliceRandomN(colds, req.ColdNum)...)
		}
	}
	if req.SoupNum > 0 {
		if len(soups) < req.SoupNum {
			resultFoods = append(resultFoods, soups...)
		} else {
			resultFoods = append(resultFoods, helper.SliceRandomN(soups, req.SoupNum)...)
		}
	}

	resp = &types.RandomResponse{}
	//返回配料字符串
	f := func(item models.Food, index int) []string {
		return strings.Split(strings.TrimPrefix(item.Material, "原料："), "、")
	}
	//flatMap将slice中每个元素转换成另一个slice，最后合并成一个slice
	flatMap := lo.FlatMap(resultFoods, f)
	//uniq函数可以去重
	resp.Materials = lo.Uniq(flatMap)

	for _, food := range resultFoods {
		resp.Foods = append(resp.Foods, food.Response())
	}

	return
}
