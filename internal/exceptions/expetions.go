package exceptions

import "errors"

var (
	ErrorFormNotFound = errors.New("Form not found")
	ErrorUserNotFound = errors.New("User not found")
	ErrorUserInvalidType = errors.New("Invalid user auth type")
	ErrorUserAuthNotFound = errors.New("User auth not found")
)
