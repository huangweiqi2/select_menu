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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.JwtTokenResponse, err error) {
	if req.Phone == "" && req.Password == "" {
		err = errs.NoArgumentErr
		return
	}

	//fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var user models.User
	err = models.DB.Where("phone=?", req.Phone).First(&user).Error
	if err != nil {
		//first方法没查到回返回ErrRecordNotFound错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errs.AuthErr
			return
		}
		return
	}

	isPasswordRight := func(userPassword, password []byte) bool {
		return bcrypt.CompareHashAndPassword(userPassword, password) == nil

	}

	if !isPasswordRight([]byte(user.Password), []byte(req.Password)) {
		err = errs.AuthErr
		return

	}
	var token string
	token, err = helper.GenerateToken(user.ID)
	if err != nil {

		return
	}

	resp = &types.JwtTokenResponse{
		AccessToken: token,
	}

	return
}
