package utilities

import "errors"

var (
	MarshallError      = errors.New("marshalling error")
	HttpCreateReqError = errors.New("new request creating error")
	HttpDoRequestError = errors.New("client do error")
	ParseRequest       = errors.New("parse request URI error")
)

type Error struct {
	Msg string
	Err error
}
