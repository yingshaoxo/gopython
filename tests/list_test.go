package main

import (
	"log"
	"testing"

	"github.com/yingshaoxo/gopython/list_tool"
)

func Test_list(t *testing.T) {
	list_ := []int{1, 2, 3}
	one := list_tool.Get_random_element_from_list(list_)
	log.Println(one)
}
