package constant

import (
	"online-book-store/internal/model"

	"github.com/gofiber/fiber/v2"
)

const (
	ErrCodeBadRequest      = 4000
	ErrCodeUnauthenticated = 4010
	ErrCodeForbidden       = 4030
	ErrCodeNotFound        = 4040
	ErrCodeConflict        = 4090
	ErrCodeUnprocessable   = 4220
	ErrCodeInternalServer  = 5000

	ErrCodeAccountNotFound   = 4041
	ErrCodeRegisteredAccount = 4221
)

const (
	ErrMessageInternalServer  = "internal server error"
	ErrMessageBadRequest      = "bad request"
	ErrMessageNotFound        = "not found"
	ErrMessageUnauthenticated = "unauthenticated"
	ErrMessageForbidden       = "forbidden"
	ErrMessageConflict        = "conflict"
	ErrMessageUnprocessable   = "unprocessable entity"

	ErrMessageAccountNotFound   = "account not found"
	ErrMessageRegisteredAccount = "account already registered"
)

var (
	ErrInternalServer = &model.AppError{
		HttpStatus: fiber.StatusInternalServerError,
		Code:       ErrCodeInternalServer,
		Message:    ErrMessageInternalServer,
	}
	ErrBadRequest = &model.AppError{
		HttpStatus: fiber.StatusBadRequest,
		Code:       ErrCodeBadRequest,
		Message:    ErrMessageBadRequest,
	}
	ErrNotFound = &model.AppError{
		HttpStatus: fiber.StatusNotFound,
		Code:       ErrCodeNotFound,
		Message:    ErrMessageNotFound,
	}
	ErrUnauthenticated = &model.AppError{
		HttpStatus: fiber.StatusUnauthorized,
		Code:       ErrCodeUnauthenticated,
		Message:    ErrMessageUnauthenticated,
	}
	ErrForbidden = &model.AppError{
		HttpStatus: fiber.StatusForbidden,
		Code:       ErrCodeForbidden,
		Message:    ErrMessageForbidden,
	}
	ErrConflict = &model.AppError{
		HttpStatus: fiber.StatusConflict,
		Code:       ErrCodeConflict,
		Message:    ErrMessageConflict,
	}
	ErrUnprocessable = &model.AppError{
		HttpStatus: fiber.StatusUnprocessableEntity,
		Code:       ErrCodeUnprocessable,
		Message:    ErrMessageUnprocessable,
	}
	ErrRegisteredAccount = &model.AppError{
		HttpStatus: fiber.StatusUnprocessableEntity,
		Code:       ErrCodeRegisteredAccount,
		Message:    ErrMessageRegisteredAccount,
	}
	ErrAccountNotFound = &model.AppError{
		HttpStatus: fiber.StatusNotFound,
		Code:       ErrCodeAccountNotFound,
		Message:    ErrMessageAccountNotFound,
	}
)
