package function_tool

import (
	"fmt"
	"reflect"
)

func Print_out_propertys_of_a_function(function interface{}) {
	functionType := reflect.TypeOf(function)
	functionValue := reflect.ValueOf(function)
	functionName := functionType.Name()
	functionTypeString := functionType.String()
	functionValueString := functionValue.String()
	functionKind := functionValue.Kind()
	functionKindString := functionValue.Type().String()
	functionNumIn := functionType.NumIn()
	functionNumOut := functionType.NumOut()
	functionIn := make([]string, functionNumIn)
	functionOut := make([]string, functionNumOut)
	for i := 0; i < functionNumIn; i++ {
		functionIn[i] = functionType.In(i).String()
	}
	for i := 0; i < functionNumOut; i++ {
		functionOut[i] = functionType.Out(i).String()
	}
	fmt.Println("functionName:", functionName)
	fmt.Println("functionTypeString:", functionTypeString)
	fmt.Println("functionValueString:", functionValueString)
	fmt.Println("functionKind:", functionKind)
	fmt.Println("functionKindString:", functionKindString)
	fmt.Println("functionNumIn:", functionNumIn)
	fmt.Println("functionNumOut:", functionNumOut)
	fmt.Println("functionIn:", functionIn)
	fmt.Println("functionOut:", functionOut)
}
