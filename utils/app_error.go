package utils

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrorCode    string
	ErrorMessage string
	ErrorType    int
}

func (e AppError) Error() string {
	return fmt.Sprintf("type: %d, code: %s, err: %s", e.ErrorType, e.ErrorCode, e.ErrorMessage)
}

// klasifikasi error ada dari http (technical) dan dari user
// transaksi saldo tidak mencukupi -> error dari user

func RequiredError() error {
	return AppError{
		ErrorCode:    "X01",
		ErrorMessage: "Input cant be empty",
		ErrorType:    http.StatusBadRequest,
	}
}

func UnauthorizedError() error {
	return AppError{
		ErrorCode:    "X04",
		ErrorMessage: "Unauthorized user",
		ErrorType:    http.StatusUnauthorized,
	}
}

func DataNotFoundError() error {
	return AppError{
		ErrorCode:    "X02",
		ErrorMessage: "No Data Found",
	}
}
