package scripts

import "fmt"

type scriptError struct {
	msg string
}

func newScriptError(msg string) scriptError {
	return scriptError{
		msg: fmt.Sprintf("script: %s", msg),
	}
}

func (err scriptError) Error() string {
	return err.msg
}
