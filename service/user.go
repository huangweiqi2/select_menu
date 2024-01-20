package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"select_menu/helper"
	"select_menu/models"
)

// Login
// @tags 公共方法
// @Summary 用户登陆
// @Param phone formData string true "13112345678"
// @Param password formData string true "password"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	if phone == "" && password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
	}

	//fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var user models.User
	err := models.DB.Where("phone=?", phone).First(&user).Error
	if err != nil {
		//first方法没查到回返回ErrRecordNotFound错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

	isPasswordRight := func(userPassword, password []byte) bool {
		return bcrypt.CompareHashAndPassword(userPassword, password) == nil

	}

	if !isPasswordRight([]byte(user.Password), []byte(password)) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return

	}
	var token string
	token, err = helper.GenerateToken(user.ID)
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
	var existUser models.User
	err := models.DB.Where("phone=?", phone).Model(new(models.User)).First(&existUser).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "重试",
			})
		}
	}
	if existUser.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该手机号已经存在",
		})
		return
	}
	//插入数据
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		Username: username,
		Password: string(fromPassword),
		Email:    email,
		Phone:    phone,
	}
	err = models.DB.Create(user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create data Error" + err.Error(),
		})
		return
	}
	//生成token
	var token string
	token, err = helper.GenerateToken(user.ID)
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
