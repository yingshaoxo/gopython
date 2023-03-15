// We handle variable related problems here, for example, is the variable nullable
package variable_tool

type Nullable[T any] struct {
	Is_null bool
	Value   T
}

type Result[T any] struct {
	value Nullable[T]
	Error string
}
