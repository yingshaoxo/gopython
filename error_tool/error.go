package error_tool

import (
	"errors"
)

func Get_error_memory_pointer_from_string(error_message string) *error {
	var err error = errors.New(error_message)
	return &err
}

// https://github.com/ez4o/go-try

type Exception struct {
	err string
}

func (e Exception) Catch(error_handle_function func(err string)) {
	error_handle_function(e.err)
}

func Try(the_function_to_execute func()) Exception {
	var exception_we_got Exception

	func() {
		defer func() {
			recover_data_we_got := recover()
			if recover_data_we_got != nil {
				// find out exactly what the error was and set err
				switch recover_data_we_got.(type) {
				case string:
					exception_we_got = Exception{err: recover_data_we_got.(string)}
				case error:
					exception_we_got = Exception{err: recover_data_we_got.(error).Error()}
				default:
					// Fallback err (per specs, error strings should be lowercase w/o punctuation
					exception_we_got = Exception{err: "unknown error"}
				}
			}
		}()

		the_function_to_execute()
	}()

	return exception_we_got
}
