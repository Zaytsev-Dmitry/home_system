package utilities

import (
	"fmt"
)

type Error struct {
	Msg string
	Err error
}

func Fail(error Error) error {
	logger := GetLogger()
	text := error.Msg + "." + " Wrapped: " + error.Err.Error()
	logger.Error(text)
	return fmt.Errorf(text)
}
