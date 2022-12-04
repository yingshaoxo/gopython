package main

import (
	"testing"

	"github.com/yingshaoxo/gopython/stringTool"
)

func Test_1_plus_1_equal_to_2(t *testing.T) {
	if (1 + 1) != 2 {
		t.Fatalf("That's wrong!")
	}
}

func Test_int_to_string(t *testing.T) {
	if stringTool.IntToString(5) != "5" {
		t.Fatalf("That's wrong!")
	}
}

// func Test_wrong_condition(t *testing.T) {
// 	t.Fatalf("I'm wrong!")
// }
