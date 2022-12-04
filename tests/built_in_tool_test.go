package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/built_in_functions"
)

func Test_print(t *testing.T) {
	built_in_functions.Print("hi")
}
