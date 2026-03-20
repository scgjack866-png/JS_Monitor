package middleware

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

const jwtKey = "adtkls"

// 创建token
func CreateToken(m map[string]string, keys ...string) string {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

// 解析token
func ParseToken(tokenString string, keys ...string) (map[string]string, bool) {
	tokenString = extractToken(tokenString)
	if tokenString == "" {
		return nil, false
	}

	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil || token == nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapData := make(map[string]string)
		for index, val := range claims {
			mapData[index] = fmt.Sprintf("%v", val)
		}
		return mapData, true
	}
	return nil, false
}

func extractToken(tokenString string) string {
	fields := strings.Fields(strings.TrimSpace(tokenString))
	if len(fields) == 0 {
		return ""
	}
	if len(fields) == 1 {
		return fields[0]
	}
	if len(fields) == 2 && strings.EqualFold(fields[0], "Bearer") {
		return fields[1]
	}
	return ""
}
