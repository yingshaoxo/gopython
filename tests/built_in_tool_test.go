package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/built_in_functions"
)

func Test_print(t *testing.T) {
	built_in_functions.Print("hi")
}

func Test_type(t *testing.T) {
	a := "abc"
	b := "abc"
	c := 35
	if built_in_functions.Type(a) == built_in_functions.Type(b) {
		built_in_functions.Print("test_1_passed")
	}
	if built_in_functions.Type(a) != built_in_functions.Type(c) {
		built_in_functions.Print("test_2_passed")
	}
	println(built_in_functions.Type(a).Name())
}
