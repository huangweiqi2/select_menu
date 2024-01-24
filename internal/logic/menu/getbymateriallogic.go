package menu

import (
	"context"
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

func NewGetByMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByMaterialLogic {
	return &GetByMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByMaterialLogic) GetByMaterial(req *types.GetByMaterialRequest) (resp *types.RandomResponse, err error) {
	material := strings.Split(req.Material, "、")
	//读取数据
	foods := make([]models.Food, 0)
	err = models.DB.Model(new(models.Food)).Find(&foods).Error
	if err != nil {
		err = errs.QueryModelErr
	}
	//计算配对率
	matchMap := make(map[string]int, len(foods))
	for _, food := range foods {
		matchMap[food.Name] = MatchRate(material, food.Response().Material)
	}
	//排序
	nameMap := sortMap(matchMap, 3)
	//返回结果
	resp = &types.RandomResponse{}
	for _, v := range nameMap {
		for _, food := range foods {
			if food.Name == v {
				resp.Foods = append(resp.Foods, food.Response())
			}
		}
	}
	return
}

// 计算配对了率
func MatchRate(a, b []string) (rate int) {
	var n = len(a) + len(b)
	m := make(map[string]struct{}, n)
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		m[v] = struct{}{}
	}
	rate = int((float64(len(m)) / float64(n)) * 100)
	return
}

// 排序函数
func sortMap(m map[string]int, n int) (resMap map[int]string) {
	invMap := make(map[int]string, len(m))
	values := make([]int, 0)
	for k, v := range m {
		invMap[v] = k
		values = append(values, v)
	}
	//int切片正序
	sort.Ints(values)
	for i, value := range values[:n] {
		resMap[i] = invMap[value]
	}
	return
}
