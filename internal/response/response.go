package response

import "encoding/json"

type Response struct {
	Status bool    `json:"status"`
	Data   any 	   `json:"data,omitempty"`
	Errors []APIError `json:"errors,omitempty"`
}

type APIError struct {
	Message string `json:"message"`
	Code 	string `json:"code"`
}

func (r Response) ToJson() string {
	d, err := json.Marshal(r)

	if err != nil {
		panic(err)
	}

	return string(d)
}

func GetUnauthorizedResponse() Response {
	return Response{
		Status: false,
		Errors: []APIError{
			{
				Message: "Unauthorized",
				Code: "401",
			},
		},
	}
}

func GetErrorResponse(message string, code string) Response  {
	return Response {
		Status: false,
		Errors: []APIError{
			{
				Message: message,
				Code: code,
			},
		},
	}
}