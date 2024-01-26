package menu

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"select_menu/internal/errs"
	"select_menu/internal/svc"
	"select_menu/internal/types"
	"select_menu/models"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type FoodRate struct {
	Food models.Food
	Rate int
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

	foodChan := make(chan models.Food)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup, foodChan chan<- models.Food, foods []models.Food) {
		defer func() {
			fmt.Println("close foodChan")
			close(foodChan)
			wg.Done()
		}()
		for _, food := range foods {
			foodChan <- food
		}
	}(&wg, foodChan, foods)

	f := func(wg *sync.WaitGroup, foodChan <-chan models.Food, materialMap map[string]struct{}) (foodRateChan chan FoodRate) {
		foodRateChan = make(chan FoodRate)
		wg.Add(1)
		go func(wg *sync.WaitGroup, foodChan <-chan models.Food, foodRateChan chan<- FoodRate, materialMap map[string]struct{}) {

			defer func() {
				fmt.Println("close foodRateChan")
				close(foodRateChan)
				wg.Done()
			}()
			for food := range foodChan {
				foodRateChan <- FoodRate{
					Rate: MatchRate2(materialMap, food.Response().Material),
					Food: food,
				}
			}
		}(wg, foodChan, foodRateChan, materialMap)
		return foodRateChan
	}

	chans := make([]<-chan FoodRate, 0, 3)
	for i := 0; i < 3; i++ {
		chans = append(chans, f(&wg, foodChan, materialMap))
	}

	foodRates := make([]FoodRate, 0, len(foods))

	for {
		select {
		case foodRate, ok := <-chans[0]:
			if !ok {
				fmt.Println("chan1 close")
				chans[0] = nil
			}
			foodRates = append(foodRates, foodRate)
		case foodRate, ok := <-chans[1]:
			if !ok {
				fmt.Println("chan2 close")
				chans[1] = nil
			}
			foodRates = append(foodRates, foodRate)
		case foodRate, ok := <-chans[2]:
			if !ok {
				fmt.Println("chan3 close")
				chans[2] = nil
			}
			foodRates = append(foodRates, foodRate)
		case <-time.After(time.Second * 2):
			break
		}
		if chans[0] == nil && chans[1] == nil && chans[2] == nil {
			break

		}

	}

	wg.Wait()

	//foodRates := make([]FoodRate, len(foods))
	//for _, food := range foods {
	//	foodRates = append(foodRates, FoodRate{
	//		Food: food,
	//		rate: MatchRate2(materialMap, food.Response().Material),
	//	})
	//}
	//排序
	sort.Slice(foodRates, func(i, j int) bool {
		return foodRates[i].Rate > foodRates[j].Rate
	})
	//返回结果
	if len(foodRates) > 5 {
		foodRates = foodRates[:5:5]
	}
	resp = &types.RandomResponse{}

	for _, rate := range foodRates {
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

func merge(inCh ...<-chan FoodRate) <-chan FoodRate {
	outCh := make(chan FoodRate, 2)
	var wg sync.WaitGroup
	for _, ch := range inCh {
		wg.Add(1)
		go func(wg *sync.WaitGroup, in <-chan FoodRate) {
			defer wg.Done()
			for val := range in {
				outCh <- val
			}
		}(&wg, ch)
	}

	// 重要注意，wg.Wait() 一定要在goroutine里运行，否则会引起deadlock
	go func() {
		wg.Wait()
		fmt.Println("close(outCh)")
		close(outCh)
	}()

	return outCh
}
