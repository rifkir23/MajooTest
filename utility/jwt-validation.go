package utility

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserValidateDTO struct {
	AccessUuid string `json:"access_uuid"`
	UserId     uint64 `json:"user_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Exp        string `json:"exp"`
}

func ValidateJwtToken(r *http.Request) (*UserValidateDTO, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid := claims["access_uuid"].(string)
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		username := claims["username"].(string)
		name := claims["name"].(string)
		email := claims["email"].(string)
		phone := claims["phone"].(string)
		address := claims["address"].(string)
		exp := claims["exp"].(float64)
		expInt := int64(exp)
		tm := time.Unix(expInt, 0)
		timeStr := tm.String()
		return &UserValidateDTO{
			AccessUuid: accessUuid,
			UserId:     userId,
			Username:   username,
			Name:       name,
			Email:      email,
			Phone:      phone,
			Address:    address,
			Exp:        timeStr,
		}, nil
	}
	return nil, err
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
