package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Username string `gorm:"column:username;type:varchar(255)" json:"username"`
	jwt.StandardClaims
}

// GetMd5
// Md5加密
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken
// 生成token
var myKey = []byte("select-menu")

func GenerateToken(username string) (string, error) {
	userClaims := &UserClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err

	}
	return tokenString, nil

}

// AnalyseToken
// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("Analyse Token Error" + err.Error())
	}
	return userClaim, nil

}
