package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

type JwtToken struct {
	Token string `json:"token"`
}

//根据 Id生成Token
func GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	return token.SignedString([]byte("secret"))
}

func GetUserIdFromToken(tokenStr string) (int, bool) {
	token, error := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte("secret"), nil
	})
	if error != nil {
		fmt.Println("error is ", error)
	}
	if !token.Valid {
		return 0, false
	}
	return GetIdFromClaims("id", token.Claims), true
}

func GetIdFromClaims(key string, claims jwt.Claims) int {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if k.Interface().(string) == key {
				return int(value.Interface().(float64))
			}
		}
	}
	return 0
}