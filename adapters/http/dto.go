package http_controllers

import "github.com/samber/mo"

type Status string

var (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type GenericAPIResponse[T any] struct {
	Status Status            `json:"status"`
	Error  mo.Option[string] `json:"error"`
	Data   T                 `json:"data"`
}

func NewSuccessResponse[T any](data T) GenericAPIResponse[T] {
	return GenericAPIResponse[T]{
		Status: StatusSuccess,
		Data:   data,
		Error:  mo.None[string](),
	}
}

func NewErrorResponse(err error) GenericAPIResponse[any] {
	return GenericAPIResponse[any]{
		Status: StatusError,
		Data:   nil,
		Error:  mo.Some(err.Error()),
	}
}

func NewErrorStringResponse(err string) GenericAPIResponse[any] {
	return GenericAPIResponse[any]{
		Status: StatusError,
		Data:   nil,
		Error:  mo.Some(err),
	}
}
