package auth

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"select_menu/helper"
	"select_menu/internal/errs"
	"select_menu/models"

	"select_menu/internal/svc"
	"select_menu/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.JwtTokenResponse, err error) {
	if req.Username == "" || req.Password == "" || req.Phone == "" {
		err = errs.NoArgumentErr
		return
	}
	//判断手机号是否存在
	var existUser models.User
	err = models.DB.Where("phone=?", req.Phone).Model(new(models.User)).First(&existUser).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	if existUser.ID > 0 {
		err = errs.PhoneExist
		return
	}
	//插入数据
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &models.User{
		Username: req.Username,
		Password: string(fromPassword),
		Email:    req.Email,
		Phone:    req.Phone,
	}
	err = models.DB.Create(user).Error
	if err != nil {
		err = errs.CreatModelErr
		return
	}
	//生成token
	var token string
	token, err = helper.GenerateToken(user.ID)
	if err != nil {
		err = errs.GenerateTokenErr
		return
	}
	resp = &types.JwtTokenResponse{
		AccessToken: token,
	}
	return
}
