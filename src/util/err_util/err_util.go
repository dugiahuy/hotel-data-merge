package err_util

import (
	"net/http"
)

type StatusCoder interface {
	StatusCode() int
}

// Error declarations
var (
	ErrNotFound       = errNotFound{}
	ErrBadRequest     = errBadRequest{}
	ErrInternalServer = errInternalServer{}
	ErrUnauthorized   = errUnauthorized{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "not found hotel"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}

type errInternalServer struct{}

func (errInternalServer) Error() string {
	return "internal server error"
}
func (errInternalServer) StatusCode() int {
	return http.StatusInternalServerError
}

type errBadRequest struct{}

func (errBadRequest) Error() string {
	return "bad request"
}
func (errBadRequest) StatusCode() int {
	return http.StatusBadRequest
}

type errUnauthorized struct{}

func (errUnauthorized) Error() string {
	return "Unauthorized, access denied"
}
func (errUnauthorized) StatusCode() int {
	return http.StatusUnauthorized
}
