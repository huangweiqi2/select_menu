package menu

import (
	"context"
	"select_menu/internal/svc"
	"select_menu/internal/types"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByMaterialLogic {
	return &GetByMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByMaterialLogic) GetByMaterial(req *types.GetByMaterialRequest) (resp *types.RandomResponse, err error) {
	//materials := strings.Split(req.Material, "、")
	//foods := make([]models.Food, 0)
	//err = models.DB.Model(new(models.Food)).Find(&foods).Error
	//if err != nil {
	//	err = errs.QueryModelErr
	//}
	//matchMap := make(map[string]int)
	//
	//for _, food := range foods {
	//
	//}
	//make()
	return
}

// 计算配对了率
func MatchRate(a, b []string) (rate int) {
	return
}

// 排序函数
func sortMap(m map[string]int, n int) (sortMap map[int]string) {
	invMap := make(map[int]string, len(m))
	values := make([]int, len(m))
	for k, v := range m {
		invMap[v] = k
		values = append(values, v)
	}
	//int切片倒序
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for i, value := range values[:n] {
		sortMap[i] = invMap[value]
	}
	return
}

// 过滤不需要的菜名
func filterName(food types.FoodResponse, materials []string) (fss types.FoodResponse) {
	//for i, material := range materials {
	//	material.
	//}
	return
}
