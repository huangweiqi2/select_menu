package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"select_menu/models"
	"strconv"
)

// Random
// @tags 公共方法
// @Summary 随机选择
// @Param status query int ture "status"
// @Param number query string ture "number"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /random [get]
func Random(c *gin.Context) {
	status, _ := strconv.Atoi(c.Query("status"))
	number := c.Query("number")
	if status == 0 || number == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "传入参数不正确",
		})
		return
	}
	data := make([]*models.Food, 0)
	err := models.DB.Model(new(models.Food)).Where("status=?", status).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Query mysql Error" + err.Error(),
		})
		return
	}

}
