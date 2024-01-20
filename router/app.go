package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "select_menu/docs"
	"select_menu/service"
)

func Router(r *gin.Engine) {
	//创建一个默认的路由引擎
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//用户路由
	r.POST("/login", service.Login)
	r.POST("/register", service.Register)

	//食物
	r.GET("/random", service.Random)

}
