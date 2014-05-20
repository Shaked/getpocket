package utils

type RequesterError interface {
	ErrorCode() int
}

type RequestError struct {
	error
	RequesterError
	xErrorCode int
	xErrorText error
}

func NewRequestError(xErrorCode int, xErrorText error) *RequestError {
	return &RequestError{
		xErrorCode: xErrorCode,
		xErrorText: xErrorText,
	}
}

func (e *RequestError) Error() string {
	return e.xErrorText.Error()
}

func (e *RequestError) ErrorCode() int {
	return e.xErrorCode
}
