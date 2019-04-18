package httperrors

import "net/http"

type Error interface {
	error
	WriteTo(http.ResponseWriter)
}

type httpError struct {
	status  int
	message string
}

func (h httpError) WriteTo(rw http.ResponseWriter) {
	http.Error(rw, h.message, h.status)
}

func (h httpError) Error() string {
	return h.message
}

func NewHttpError(m string, status int) Error {
	return httpError{status, m}
}

func NewBadRequestError(m string) Error {
	return NewHttpError(m, http.StatusBadRequest)
}

func NewUnauthorizedError(m string) Error {
	return NewHttpError(m, http.StatusUnauthorized)
}

func NewInternalServerError(m string) Error {
	return NewHttpError(m, http.StatusInternalServerError)
}
