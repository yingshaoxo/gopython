package jwt_tool

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yingshaoxo/gopython/string_tool"
)

func Jwt_encode(data map[string]interface{}, secret_string string) string {
	new_data := jwt.MapClaims{}
	for key, value := range data {
		new_data[key] = value
	}

	secret_string_bytes := string_tool.String_to_bytes(secret_string)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, new_data)
	token_string, _ := token.SignedString(secret_string_bytes)

	return token_string
}

func Jwt_decode(jwt_string string, secret_string string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwt_string, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret_string), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data := make(map[string]interface{})
		for key, value := range claims {
			data[key] = value
		}
		return data, nil
	} else {
		return nil, fmt.Errorf("Token error.")
	}
}
