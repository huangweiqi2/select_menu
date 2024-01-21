package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
)

type UserClaims struct {
	ID uint `json:"id"`
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

func GenerateToken(ID uint) (string, error) {
	userClaims := &UserClaims{
		ID:             ID,
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

// 16位随机字符串
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// 密钥，实际应用中应该从环境变量或文件中获取
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

// Base64编码和解码方法
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// 加密方法
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return Encode(cipherText), nil
}

// 解密方法
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}

	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func SliceRandomN[T any](slice []T, n int) []T {
	// 先将切片进行打乱
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	// 返回前n个元素
	return slice[:n]
}
