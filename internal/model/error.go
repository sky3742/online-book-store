package model

type AppError struct {
	HttpStatus int
	Code       int
	Message    string
}
