package main

import (
	"log"
	"testing"

	"github.com/yingshaoxo/gopython/jwt_tool"
)

func Test_jwt(t *testing.T) {
	secret := "no way"

	data := make(map[string]interface{})
	data["user"] = "yingshaoxo"

	jwt_string := jwt_tool.Jwt_encode(data, secret)
	log.Println(jwt_string)

	new_data, _ := jwt_tool.Jwt_decode(jwt_string, secret)

	log.Println(data)
	log.Println(new_data)
}
