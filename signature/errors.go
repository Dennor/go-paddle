package signature

type VerificationError struct {
	error
}

func NewVerificationError(err error) VerificationError {
	return VerificationError{err}
}
