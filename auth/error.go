package auth

type AuthenticatorError interface {
	ErrorCode() int
}

type AuthError struct {
	error
	AuthenticatorError
	xErrorCode int
	xErrorText error
}

func NewAuthError(xErrorCode int, xErrorText error) *AuthError {
	return &AuthError{
		xErrorCode: xErrorCode,
		xErrorText: xErrorText,
	}
}

func (e *AuthError) Error() string {
	return e.xErrorText.Error()
}

func (e *AuthError) ErrorCode() int {
	return e.xErrorCode
}
