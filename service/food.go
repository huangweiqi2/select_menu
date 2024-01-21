package service

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"log"
	"net/http"
	"select_menu/helper"
	"select_menu/models"
	"select_menu/router"
)

// Random
// @tags 公共方法
// @Summary 随机选择
// @Param status query int ture "status"
// @Param number query string ture "number"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /random [get]
func Random(c *gin.Context) {
	var req router.RandomReq
	if c.ShouldBind(&req) == nil {
		log.Println("====== Only Bind By Query String ======")
	}

	if err := req.Valid(); err != nil {
		return
	}
	resp, err := RandomLogic(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Query mysql Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": resp,
	})
}

func RandomLogic(req router.RandomReq) (resp router.RandomResp, err error) {
	var foods []models.Food
	err = models.DB.Model(new(models.Food)).Find(&foods).Error
	if err != nil {
		return
	}
	hots := make([]router.FoodResp, 0, len(foods))
	colds := make([]router.FoodResp, 0, len(foods))
	soups := make([]router.FoodResp, 0, len(foods))

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

	f := func(item router.FoodResp, index int) []string {
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
