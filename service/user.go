package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"select_menu/helper"
	"select_menu/models"
)

// Login
// @tags 公共方法
// @Summary 用户登陆
// @Param username formData string true "name"
// @Param password formData string true "password"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" && password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
	}
	password = helper.GetMd5(password)
	data := new(models.User)
	err := models.DB.Where("username=? AND password=?", username, password).First(&data).Error
	if err != nil {
		//first方法没查到回返回ErrRecordNotFound错误
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error :" + err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Generate token Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// Register
// @tags 公共方法
// @Summary 用户注册
// @Param username formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string true "phone"
// @Param email formData string false "email"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	if username == "" || password == "" || phone == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数传入不正确",
		})
	}
	//判断手机号是否存在
	var cnt int64
	err := models.DB.Where("phone=?", phone).Model(new(models.User)).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get user Error" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该手机号已经存在",
		})
		return
	}
	//插入数据
	data := &models.User{
		Username: username,
		Password: helper.GetMd5(password),
		Email:    email,
		Phone:    phone,
	}
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create data Error" + err.Error(),
		})
		return
	}
	//生成token
	token, err := helper.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Generate token Error" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
			"msg":   "注册成功",
		},
	})
	return

}
