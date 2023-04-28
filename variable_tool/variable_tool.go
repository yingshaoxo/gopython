// We handle variable related problems here, for example, is the variable nullable
package variable_tool

import (
	"reflect"
	"strings"
)

// T here means a general type
type Type_Nullable[T any] struct {
	Is_null bool
	Value   T
}

func (self Type_Nullable[T]) Set_to_null() Type_Nullable[T] {
	var item T = self.Value
	return Type_Nullable[T]{
		Value:   item,
		Is_null: true,
	}
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

func Is_the_variable_a_list_object(a_variable any) bool {
	object_type_representation := reflect.TypeOf(a_variable).Kind()
	if (object_type_representation != reflect.Array) && (object_type_representation != reflect.Slice) {
		return false
	} else {
		return true
	}
}

func Is_the_variable_a_dict_object(a_variable any) bool {
	object_type_representation := reflect.TypeOf(a_variable).Kind()
	if object_type_representation != reflect.Map {
		return false
	} else {
		return true
	}
}

func Is_nullable_variable(a_variable any) bool {
	var the_type = Get_variable_type_string_representation(a_variable)
	if strings.Contains(the_type, "Type_Nullable[") {
		return true
	} else {
		return false
	}
}

func Is_it_null(a_variable any) bool {
	return Get_value_from_struct_object_by_name(a_variable, "Is_null").(bool)
}

func Get_value_from_nullable_variable(a_variable any) any {
	if Is_nullable_variable(a_variable) {
		return Get_value_from_struct_object_by_name(a_variable, "Value")
	} else {
		return a_variable
	}
}

func Get_variable_type_string_representation(a_variable any) string {
	a_reflect := reflect.TypeOf(a_variable)
	name := a_reflect.Name()
	if len(strings.TrimSpace(name)) != 0 {
		return name
	} else {
		return a_reflect.String()
		// return fmt.Sprintf("%v", reflect.TypeOf(a_variable).Kind().String())
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

func Get_key_value_dict_from_struct_object(an_object_instance any, lowercase_key bool) map[string]any {
	var key_value_dict = make(map[string]any)

	object_key_representation := reflect.TypeOf(an_object_instance)
	object_value_representation := reflect.ValueOf(an_object_instance)

	for i := 0; i < object_value_representation.NumField(); i++ {
		var the_key = object_key_representation.Field(i).Name
		if lowercase_key == true {
			the_key = strings.ToLower(the_key)
		}
		// var the_type = object_type_representation.Field(i).Type.Name()
		var the_value = object_value_representation.Field(i).Interface()

		key_value_dict[the_key] = the_value
	}

	return key_value_dict
}

func Check_if_key_in_struct_object(an_object_instance any, key string, lowercase_key bool) bool {
	object_key_representation := reflect.TypeOf(an_object_instance)

	for i := 0; i < object_key_representation.NumField(); i++ {
		var the_key = object_key_representation.Field(i).Name
		if lowercase_key == true {
			the_key = strings.ToLower(the_key)
		}

		if the_key == strings.ToLower(key) {
			return true
		}
	}

	return false
}

func Get_element_type_of_a_list(a_list any) reflect.Type {
	return reflect.TypeOf(a_list).Elem()
}

func Get_a_default_value_from_reflect_type(reflect_type reflect.Type) any {
	return reflect.Zero(reflect_type).Interface()
}

func Call_struct_object_function(an_object any, method_name string, input_arguments_list []any) []any {
	// if the object you give is from reflect.ValueOf(), you need to call reflect.ValueOf().Interface()
	var type_of_an_element_in_a_list = reflect.TypeOf(an_object)
	var instance_of_a_type = reflect.Zero(type_of_an_element_in_a_list)

	var arguments []reflect.Value
	for index := 0; index < len(input_arguments_list); index++ {
		arguments = append(arguments, reflect.ValueOf(input_arguments_list[index]))
	}
	var reflect_type_output_list = instance_of_a_type.MethodByName(method_name).Call(arguments)

	var outputs []any
	for index := 0; index < len(reflect_type_output_list); index++ {
		outputs = append(outputs, reflect_type_output_list[index].Interface())
	}

	return outputs
}

func is_the_variable_an_enum_class(a_variable any) bool {
	var yes = false

	object_key_representation := reflect.TypeOf(a_variable)
	object_value_representation := reflect.ValueOf(a_variable)

	for i := 0; i < object_value_representation.NumField(); i++ {
		var the_key = object_key_representation.Field(i).Name

		if the_key == "Enum_value_" {
			yes = true
			break
		}
	}

	return yes
}

// func Convert_nullable_struct_into_dict(an_object_instance any, lowercase_the_key bool) any {
// 	if an_object_instance == nil {
// 		return nil
// 	}

// 	if Is_the_variable_a_list_object(an_object_instance) {
// 		var new_list = make([]any, 0)
// 		switch t := an_object_instance.(type) {
// 		case []any:
// 			for _, value := range t {
// 				new_list = append(new_list, Convert_nullable_struct_into_dict(value, lowercase_the_key))
// 			}
// 		}
// 		return new_list
// 	}

// 	if !Is_the_variable_a_struct_object(an_object_instance) {
// 		return an_object_instance
// 	}

// 	if is_the_variable_an_enum_class(an_object_instance) {
// 		var new_object = Get_value_from_struct_object_by_name(an_object_instance, "Enum_value_")

// 		if Is_it_null(new_object) {
// 			return nil
// 		} else {
// 			var new_value = Get_value_from_nullable_variable(new_object).(string)
// 			if Check_if_key_in_struct_object(an_object_instance, new_value, true) {
// 				if lowercase_the_key == true {
// 					return strings.ToLower(new_value)
// 				} else {
// 					return new_value
// 				}
// 			} else {
// 				return nil
// 			}
// 		}
// 	}

// 	var new_dict = make(map[string]any)

// 	object_key_representation := reflect.TypeOf(an_object_instance)
// 	object_value_representation := reflect.ValueOf(an_object_instance)

// 	types := make([]any, object_key_representation.NumField())
// 	values := make([]interface{}, object_value_representation.NumField())
// 	for i := 0; i < object_value_representation.NumField(); i++ {
// 		var the_key = object_key_representation.Field(i).Name
// 		var the_type = object_key_representation.Field(i).Type.Name()
// 		var the_value = object_value_representation.Field(i).Interface()
// 		types[i] = the_type
// 		values[i] = the_value

// 		var is_nullable bool = false
// 		if strings.Contains(the_type, "Type_Nullable[") {
// 			is_nullable = true
// 		}

// 		var new_object any = nil
// 		if is_nullable {
// 			if Is_it_null(the_value) {
// 				new_object = nil
// 			} else {
// 				var new_value = Get_value_from_nullable_variable(the_value)
// 				new_object = Convert_nullable_struct_into_dict(new_value, lowercase_the_key)
// 			}
// 		} else {
// 			new_object = Convert_nullable_struct_into_dict(the_value, lowercase_the_key)
// 		}

// 		if lowercase_the_key == true {
// 			new_dict[strings.ToLower(the_key)] = new_object
// 		} else {
// 			new_dict[the_key] = new_object
// 		}
// 	}

// 	return new_dict
// }
