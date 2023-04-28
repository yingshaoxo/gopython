// We handle variable related problems here, for example, is the variable nullable
package variable_tool

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/yingshaoxo/gopython/dict_tool"
	"github.com/yingshaoxo/gopython/json_tool"
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

/*
type User_Status struct {
	Enum_value_ variable_tool.Type_Nullable[string]
	Online      variable_tool.Type_Nullable[string]
	Offline     variable_tool.Type_Nullable[string]
}

func (self User_Status) New(value variable_tool.Type_Nullable[string]) User_Status {
	var item = User_Status{}
	item.Online = variable_tool.Nullable("Online")
	item.Offline = variable_tool.Nullable("Offline")
	item.Enum_value_ = value
	return item
}

func (self User_Status) To_dict() string {
	return self.Enum_value_.Value
}

func (self User_Status) From_dict(value string) User_Status {
	var item = User_Status{}.New(variable_tool.Nullable(value))
	return item
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
*/

// var Null_value_identify_symbol string = "It's fucking null. The stupid golang doesn't support null value in map structure, which sucks!!! And golang doesn't support lower-case exported function name, bad. And golang doesn't support disableing the unused variable warning, which is super bad!"
var Null_value_identify_symbol *string = nil

func Convert_dirty_map_into_pure_map(an_object_instance any) any {
	if an_object_instance == nil {
		return nil
	}

	if Is_the_variable_a_struct_object(an_object_instance) {
		a_map, _ := json_tool.Convert_struct_object_to_map(an_object_instance)
		return Convert_dirty_map_into_pure_map(a_map)
	}

	if Is_the_variable_a_list_object(an_object_instance) {
		var new_list = make([]any, 0)
		switch t := an_object_instance.(type) {
		case []any:
			for _, value := range t {
				new_list = append(new_list, Convert_dirty_map_into_pure_map(value))
			}
		}
		return new_list
	}

	if Is_the_variable_a_dict_object(an_object_instance) {
		var new_dict = make(map[string]any)
		for key, value := range an_object_instance.(map[string]any) {
			new_dict[key] = Convert_dirty_map_into_pure_map(value)
		}
		return new_dict
	}

	return an_object_instance
}

func Convert_nullable_struct_into_dict(an_object_instance any, lowercase_the_key bool) any {
	if an_object_instance == nil {
		return nil
	}

	if Is_the_variable_a_struct_object(an_object_instance) {
		a_map, _ := json_tool.Convert_struct_object_to_map(an_object_instance)
		return Convert_nullable_struct_into_dict(a_map, lowercase_the_key)
	}

	if Is_the_variable_a_list_object(an_object_instance) {
		var new_list = make([]any, 0)
		switch t := an_object_instance.(type) {
		case []any:
			for _, value := range t {
				if Is_the_variable_a_dict_object(value) {
					if dict_tool.Check_if_a_key_is_in_the_dict(value.(map[string]any), "Is_null") {
						if dict_tool.Get_dict_value_by_giving_a_key(value.(map[string]any), "Is_null") == true {
							new_list = append(new_list, Null_value_identify_symbol)
						} else {
							new_list = append(new_list, dict_tool.Get_dict_value_by_giving_a_key(value.(map[string]any), "Value"))
						}
						continue
					}
				}
				new_list = append(new_list, Convert_nullable_struct_into_dict(value, lowercase_the_key))
			}
		}
		return new_list
	}

	if Is_the_variable_a_dict_object(an_object_instance) {
		var new_dict = make(map[string]any)
		for key, value := range an_object_instance.(map[string]any) {
			new_key := key
			if lowercase_the_key == true {
				new_key = strings.ToLower(key)
			}

			if Is_the_variable_a_dict_object(value) {
				if dict_tool.Check_if_a_key_is_in_the_dict(value.(map[string]any), "Is_null") {
					if dict_tool.Get_dict_value_by_giving_a_key(value.(map[string]any), "Is_null") == true {
						new_dict[new_key] = Convert_nullable_struct_into_dict(Null_value_identify_symbol, lowercase_the_key)
					} else {
						new_dict[new_key] = Convert_nullable_struct_into_dict(
							dict_tool.Get_dict_value_by_giving_a_key(value.(map[string]any), "Value"), lowercase_the_key,
						)
					}
					continue
				}
			}

			new_dict[new_key] = Convert_nullable_struct_into_dict(value, lowercase_the_key)
		}

		return new_dict
	}

	return an_object_instance
}

func _convert_nullable_dict_into_compact_json_string(an_object_instance any) any {
	if an_object_instance == nil {
		return `null`
	}

	if Is_the_variable_a_struct_object(an_object_instance) {
		return `"` + Get_variable_type_string_representation(an_object_instance) + `"`
	}

	if Is_the_variable_a_list_object(an_object_instance) {
		var new_list_text string = `[`
		switch t := an_object_instance.(type) {
		case []any:
			for index, value := range t {
				new_list_text += _convert_nullable_dict_into_compact_json_string(value).(string)
				if index != len(t)-1 {
					new_list_text += `, `
				}
			}
		}
		new_list_text += `]`
		return new_list_text
	}

	if Is_the_variable_a_dict_object(an_object_instance) {
		var new_dict_string = `{`
		var index int = 0
		for key, value := range an_object_instance.(map[string]any) {
			new_dict_string += `"` + key + `"` + `: `
			new_dict_string += _convert_nullable_dict_into_compact_json_string(value).(string)
			if index != len(an_object_instance.(map[string]any))-1 {
				new_dict_string += `, `
			}
			index += 1
		}
		new_dict_string += `}`
		return new_dict_string
	}

	if (an_object_instance == Null_value_identify_symbol) || (an_object_instance == nil) {
		return `null`
	} else if Get_variable_type_string_representation(an_object_instance) == "string" {
		encoded_string, _ := json.Marshal(an_object_instance)
		return fmt.Sprintf(`"%v"`, encoded_string)
	} else {
		return fmt.Sprintf(`%v`, an_object_instance)
	}
}

func _convert_nullable_dict_into_json_string_with_indent_levels(an_object_instance any, level int) (any, int) {
	var indent string = strings.Repeat("    ", level+1)
	var previous_indent string = strings.Repeat("    ", level)

	if an_object_instance == nil {
		return `null`, level
	}

	if Is_the_variable_a_struct_object(an_object_instance) {
		return `"` + Get_variable_type_string_representation(an_object_instance) + `"`, level
	}

	if Is_the_variable_a_list_object(an_object_instance) {
		var new_list_text string = `[` + "\n"
		switch t := an_object_instance.(type) {
		case []any:
			for index, value := range t {
				child_string, _ := _convert_nullable_dict_into_json_string_with_indent_levels(value, level+1)
				if child_string.(string) != `null` {
					new_list_text += indent + child_string.(string)
					if index != len(t)-1 {
						new_list_text += `, ` + "\n"
					}
				}
			}
		}
		new_list_text += "\n" + previous_indent + `]`
		return new_list_text, level
	}

	if Is_the_variable_a_dict_object(an_object_instance) {
		var new_dict_string = `{` + "\n"
		var index int = 0
		for key, value := range an_object_instance.(map[string]any) {
			new_dict_string += indent + `"` + key + `"` + `: `
			child_string, _ := _convert_nullable_dict_into_json_string_with_indent_levels(value, level+1)
			new_dict_string += child_string.(string)
			if index != len(an_object_instance.(map[string]any))-1 {
				new_dict_string += `, ` + "\n"
			}
			index += 1
		}
		new_dict_string += "\n" + previous_indent + `}`
		return new_dict_string, level
	}

	if (an_object_instance == Null_value_identify_symbol) || (an_object_instance == nil) {
		return `null`, level
	} else if Get_variable_type_string_representation(an_object_instance) == "string" {
		encoded_string, _ := json.Marshal(an_object_instance)
		return fmt.Sprintf(`"%v"`, encoded_string), level
	} else {
		return fmt.Sprintf(`%v`, an_object_instance), level
	}
}

func Convert_nullable_struct_into_json_string(an_object_instance any) string {
	a_dict := Convert_nullable_struct_into_dict(an_object_instance, true)
	result, _ := _convert_nullable_dict_into_json_string_with_indent_levels(a_dict, 0)
	return result.(string)
}

func _convert_dict_into_nullable_dict(a_dict any, a_refrence_object_instance any) any {
	if a_dict == nil {
		return nil
	}
	if a_refrence_object_instance == nil {
		return nil
	}

	if Is_the_variable_a_list_object(a_dict) {
		var new_list = make([]any, 0)
		switch t := a_dict.(type) {
		case []any:
			for _, value := range t {
				// get child type or child object from empty list
				var type_of_an_element_in_a_list = reflect.TypeOf(a_refrence_object_instance).Elem()
				var instance_of_a_type = reflect.Zero(type_of_an_element_in_a_list).Interface()
				new_list = append(new_list, _convert_dict_into_nullable_dict(value, instance_of_a_type))
			}
		}
		return new_list
	}

	if !Is_the_variable_a_struct_object(a_refrence_object_instance) {
		// return nil if the basic element inside the tree is a dict than string, int, bool...
		if Is_the_variable_a_dict_object(a_dict) {
			return a_refrence_object_instance
		} else {
			return a_dict
		}
	}

	/*
		if is_the_variable_an_enum_class(a_refrence_object_instance) {
			return variable_tool.Nullable("").Set_to_null()
			var real_value = dict_tool.Get_dict_value_by_giving_a_key(a_dict.(map[string]any), "Enum_value_")
			if variable_tool.Get_variable_type_string_representation(real_value) == "string" {
				return variable_tool.Nullable(real_value).Set_to_null()
			} else {
				if real_value == nil {
					return variable_tool.Nullable("").Set_to_null()
				} else {
					var real_value2 = dict_tool.Get_dict_value_by_giving_a_key(real_value.(map[string]any), "Enum_value_")
					if real_value2 == nil {
						var result = Call_struct_object_function(a_refrence_object_instance, "New", []any{
							variable_tool.Nullable(
								a_refrence_object_instance,
							),
						})[0]
						return result
					} else {
						var result = Call_struct_object_function(a_refrence_object_instance, "New", []any{
							variable_tool.Nullable(
								real_value2.(string),
							),
						})[0]
						return result
					}
				}
			}
		}
	*/

	var new_dict = make(map[string]any)

	object_key_representation := reflect.TypeOf(a_refrence_object_instance)
	object_value_representation := reflect.ValueOf(a_refrence_object_instance)

	types := make([]any, object_key_representation.NumField())
	values := make([]interface{}, object_value_representation.NumField())
	for i := 0; i < object_value_representation.NumField(); i++ {
		var the_key = object_key_representation.Field(i).Name
		runes := []rune(the_key)
		runes[0] = unicode.ToLower(runes[0])
		the_key = string(runes)
		var the_type = object_key_representation.Field(i).Type.Name()
		var the_reference_value = object_value_representation.Field(i).Interface()
		types[i] = the_type
		values[i] = the_reference_value

		var is_nullable bool = false
		if strings.Contains(the_type, "Type_Nullable[") {
			is_nullable = true
		}

		var new_object any = nil
		var new_value = dict_tool.Get_dict_value_by_giving_a_key(a_dict.(map[string]any), the_key)
		if is_nullable {
			the_reference_value = Get_value_from_nullable_variable(the_reference_value)

			if new_value == nil || new_value == Null_value_identify_symbol {
				new_object = Nullable(
					the_reference_value,
				).Set_to_null()
			} else {
				// if dict_tool.Check_if_a_key_is_in_the_dict(new_value.(map[string]any), "Is_null") && dict_tool.Check_if_a_key_is_in_the_dict(new_value.(map[string]any), "Value") {
				// 	new_value = dict_tool.Get_dict_value_by_giving_a_key(new_value.(map[string]any), "Value")
				// }
				new_object = Nullable(
					_convert_dict_into_nullable_dict(
						new_value,
						the_reference_value,
					),
				)
			}
		} else {
			new_object = _convert_dict_into_nullable_dict(new_value, the_reference_value)
		}

		new_dict[the_key] = new_object
	}

	new_dict, _ = json_tool.Convert_struct_object_to_map(new_dict)
	return new_dict
}

// func replace_null_to_null_identify_symbol(json_string string) string {
// 	// pattern := regexp.MustCompile("(?P<before>\\\"\\w+?\\\"\\s*:\\s*)(?P<null>null)(?P<after>,?)")
// 	// template := `${before}"` + Null_value_identify_symbol + `"${after}`
// 	// replaced := pattern.ReplaceAllString(json_string, template)
// 	// return replaced
// }

// \"\w+?\"\s*:\s*\[(\s*null,?\s*)+\]
/*
This function can't handle the null in list

"Option": [
	null,
	null
]
*/
func Convert_json_string_into_nullable_struct[T any](json_string string, an_object_instance *T) {
	a_dict, _ := json_tool.Convert_json_string_to_map(json_string)
	new_dict := _convert_dict_into_nullable_dict(a_dict, *an_object_instance)
	json_tool.Convert_map_to_struct_object(new_dict, an_object_instance)
}
