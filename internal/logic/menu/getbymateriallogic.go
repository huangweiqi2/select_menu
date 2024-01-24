package menu

import (
	"context"
	"github.com/samber/lo"
	"select_menu/internal/errs"
	"select_menu/internal/svc"
	"select_menu/internal/types"
	"select_menu/models"
	"sort"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type FoodRate struct {
	Food models.Food
	rate int
}

func NewGetByMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByMaterialLogic {
	return &GetByMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByMaterialLogic) GetByMaterial(req *types.GetByMaterialRequest) (resp *types.RandomResponse, err error) {
	material := strings.Split(req.Material, "、")
	materialMap := lo.SliceToMap(material, func(item string) (string, struct{}) {
		return item, struct{}{}
	})
	//读取数据
	foods := make([]models.Food, 0)
	err = models.DB.Model(new(models.Food)).Find(&foods).Error
	if err != nil {
		err = errs.QueryModelErr
	}
	//计算配对率
	foodRate := make([]FoodRate, len(foods))
	for _, food := range foods {
		foodRate = append(foodRate, FoodRate{
			Food: food,
			rate: MatchRate2(materialMap, food.Response().Material),
		})
	}
	//排序
	sort.Slice(foodRate, func(i, j int) bool {
		return foodRate[i].rate < foodRate[j].rate
	})
	//返回结果
	if len(foodRate) > 5 {
		foodRate = foodRate[:5:5]
	}
	resp = &types.RandomResponse{}

	for _, rate := range foodRate {
		response := rate.Food.Response()
		resp.Foods = append(resp.Foods, response)
		resp.Materials = append(resp.Materials, response.Material...)
	}

	resp.Materials = lo.Uniq(resp.Materials)

	return
}

func MatchRate2(materialMap map[string]struct{}, b []string) (rate int) {
	var count int
	for _, v := range b {
		if _, ok := materialMap[v]; ok {
			count++
		}
	}

	rate = int((float64(count) / float64(len(b))) * 100)
	return
}
