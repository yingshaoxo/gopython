package error_tool

import (
	"errors"
)

func Get_error_memory_pointer_from_string(error_message string) *error {
	var err error = errors.New(error_message)
	return &err
}
