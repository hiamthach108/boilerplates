package errorx

import "fmt"

type AppError struct {
	Code    AppErrCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func Wrap(code AppErrCode, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: GetErrorMessage(int(code)),
		Err:     err,
	}
}

func New(code AppErrCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func GetCode(err error) AppErrCode {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code
	}
	return ErrInternal
}
