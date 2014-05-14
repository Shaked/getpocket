package auth

type AuthError struct {
	xErrorCode int
	xErrorText string
}

func NewAuthError(xErrorCode int, xErrorText string) *AuthError {
	return &AuthError{
		xErrorCode: xErrorCode,
		xErrorText: xErrorText,
	}
}

func (e *AuthError) Error() string {
	return e.xErrorText
}

func (e *AuthError) ErrorCode() int {
	return e.xErrorCode
}
