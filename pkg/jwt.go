package pkg

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secret = "chaojiwudizhuzhux"

type MyClaims struct {
	jwt.StandardClaims // 嵌入标准声明

	// 添加自定义字段
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
}

func GenToken(username string, user_id int64) (string, error) {
	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "WhiteFlowerGirl",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		Username: username,
		UserID:   user_id,
	})

	// 生成 token 字符串
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// 检查解析是否成功
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// 检查令牌是否有效
	if !token.Valid {
		fmt.Println("Invalid token")
		return nil, err
	}

	// 提取声明信息
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		fmt.Println("Failed to parse claims")
		return nil, errors.New("Failed to parse claims")
	}
	return claims, nil
}
