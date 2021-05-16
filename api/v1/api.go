package v1

import (
	"acgfate/serializer"
)

func ErrorResponse(code int, err error) serializer.Response {
	return serializer.Response{
		Code: code,
		Data: nil,
		Msg:  err.Error(),
	}
}
