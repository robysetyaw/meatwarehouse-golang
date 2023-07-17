package apperror

import (
	"fmt"
	"strconv"
	"strings"
)

type AppError struct {
	ErrorMassage string
	ErrorCode    int
}

func (ae AppError) Error() string {
	return fmt.Sprintf("%v - %v", ae.ErrorCode, ae.ErrorMassage)
}

func ErrCode(input string) int  {
	errcode, _ :=strconv.Atoi(strings.Split(input, "-")[0])
	return errcode
}

func ErrMessage(input string)string {
	return strings.Split(input, "-")[1]
}