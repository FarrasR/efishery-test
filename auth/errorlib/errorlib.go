package errorlib

var (
	ErrSomethingWrong        = CustomError{msg: SomethingWrongMessage}
	ErrPhoneNumberRegistered = CustomError{msg: PhoneNumberErrorMessage}
	ErrLogin                 = CustomError{msg: LoginErrorMessage}
	ErrInvalidJwt            = CustomError{msg: InvalidJwtErrorMessage}
)

const (
	SomethingWrongMessage   = "Something went wrong"
	PhoneNumberErrorMessage = "Phone number already registered"
	LoginErrorMessage       = "Phone number doesn't found or wrong password"
	InvalidJwtErrorMessage  = "JWT is invalid"
)

type CustomError struct{ msg string }

func (ce CustomError) Error() string {
	return ce.msg
}
