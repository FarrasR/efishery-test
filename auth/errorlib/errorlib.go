package errorlib

var (
	ErrPhoneNumberRegistered = CustomError{msg: PhoneNumberError}
	ErrSomethingWrong        = CustomError{msg: SomethingWrong}
)

const (
	SomethingWrong   = "Somethinmg is wrong"
	PhoneNumberError = "Phone number already registered"
)

type CustomError struct{ msg string }

func (nfe CustomError) Error() string {
	return nfe.msg
}
