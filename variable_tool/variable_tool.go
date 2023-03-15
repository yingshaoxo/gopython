// We handle variable related problems here, for example, is the variable nullable
package variable_tool

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
