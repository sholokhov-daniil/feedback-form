package response

import "encoding/json"

type ErrorCode string

const (
	ErrorCodeServer string = "SERVER_ERROR"
	ErrorCodeAuthError string = "AUTH_ERROR"
	ErrorCodeNotFound string = "NOT_FOUND"
)

type Error struct {
	Message string `json:"message"`
	Code 	string `json:"code"`
}

func (e Error) ToJson() string {
	d, err := json.Marshal(e)

	if err != nil {
		panic(err)
	}

	return string(d)

}

func CreateUnauthorizedResponse() Error {
	return Error{
		Message: "Unauthorized",
			Code: ErrorCodeAuthError,
	}
}

func CreateServerErrorResponse(message string) Error {
	return Error{
		Message: message,
		Code: ErrorCodeServer,
	}
}

func CreateNotFoundErrorResponse(message string) Error {
	return Error{
		Message: message,
		Code: ErrorCodeNotFound,
	}
}