package response

import (
	"encoding/json"
)

type ErrorCode string

const (
	ErrorCodeServer string = "SERVER_ERROR"
	ErrorCodeAuthError string = "AUTH_ERROR"
)

type Response struct {
	Status bool    `json:"status"`
	Data   any 	   `json:"data"`
	Errors []Error `json:"errors"`
}

type Error struct {
	Message string `json:"message"`
	Code 	string `json:"code"`
}

func New(d any) Response {
	return Response{
		Status: true,
		Data: d,
		Errors: []Error{},
	}
}

func (r Response) ToJson() string {
	d, err := json.Marshal(r)

	if err != nil {
		panic(err)
	}

	return string(d)
}

func CreateUnauthorizedResponse() Response {
	return Response{
		Status: false,
		Errors: []Error{
			{
				Message: "Unauthorized",
				Code: ErrorCodeAuthError,
			},
		},
	}
}

func CreateServerErrorResponse(message string) Response {
	return CreateErrorResponse(message, ErrorCodeServer)
}

func CreateErrorResponse(message string, code string) Response  {
	return Response {
		Status: false,
		Errors: []Error{
			{
				Message: message,
				Code: code,
			},
		},
	}
}