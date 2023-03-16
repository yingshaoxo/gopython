// We handle variable related problems here, for example, is the variable nullable
package variable_tool

import "reflect"

// T here means a general type
type Type_Nullable[T any] struct {
	Is_null bool
	Value   T
}

func Nullable[T any](value T) Type_Nullable[T] {
	var item T = value
	return Type_Nullable[T]{
		Value:   item,
		Is_null: false,
	}
}

type _Result[T any] struct {
	Value Type_Nullable[T]
	Error string
}

func Result[T any](value T) _Result[T] {
	var item T = value
	return _Result[T]{
		Value: Nullable[T](item),
		Error: "",
	}
}

func Is_the_variable_a_struct_object(a_variable any) bool {
	object_type_representation := reflect.TypeOf(a_variable).Kind()
	if object_type_representation != reflect.Struct {
		return false
	} else {
		return true
	}
}

func Get_value_from_struct_object_by_name(an_object_instance any, key string) any {
	object_key_representation := reflect.TypeOf(an_object_instance)
	object_value_representation := reflect.ValueOf(an_object_instance)

	for i := 0; i < object_value_representation.NumField(); i++ {
		var the_key = object_key_representation.Field(i).Name
		// var the_type = object_type_representation.Field(i).Type.Name()
		var the_value = object_value_representation.Field(i).Interface()

		if the_key == key {
			Is_the_variable_a_struct_object(the_value)
			return the_value
		}
	}

	return nil
}
