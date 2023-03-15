package variable_tool

type Nullable[T any] struct {
	Is_null bool
	Value   T
}
