package built_in_functions

import (
	"fmt"
	"reflect"
)

func Print(anything interface{}) {
	fmt.Println(anything)
}

func Type(anything interface{}) reflect.Type {
	return reflect.TypeOf(anything)
}

// A
// abs()
// aiter()
// all()
// any()
// anext()
// ascii()

// B
// bin()
// bool()
// breakpoint()
// bytearray()
// bytes()

// C
// callable()
// chr()
// classmethod()
// compile()
// complex()

// D
// delattr()
// dict()
// dir()
// divmod()

// E
// enumerate()
// eval()
// exec()

// F
// filter()
// float()
// format()
// frozenset()

// G
// getattr()
// globals()

// H
// hasattr()
// hash()
// help()
// hex()

// I
// id()
// input()
// int()
// isinstance()
// issubclass()
// iter()
// L
// len()
// list()
// locals()

// M
// map()
// max()
// memoryview()
// min()

// N
// next()

// O
// object()
// oct()
// open()
// ord()

// P
// pow()
// print()
// property()

// R
// range()
// repr()
// reversed()
// round()

// S
// set()
// setattr()
// slice()
// sorted()
// staticmethod()
// str()
// sum()
// super()

// T
// tuple()
// type()

// V
// vars()

// Z
// zip()

// _
// __import__()
