package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/yingshaoxo/gopython/error_tool"
)

func Test_error_tool(t *testing.T) {
	error_tool.Try(func() {
		i := 3
		for i != -1 {
			result := 100 / i
			fmt.Println(result, i)

			i -= 1
		}
	}).Catch(func(err string) {
		log.Println(err)
	})
}
