package errors

import "net/http"

// CREATE
//  -> Ok
//  -> Invalid Request
//  -> Conflict
//  -> Unauthorized
//  -> Something Else

// QUERY
//  -> Ok
//  -> Not Found
//  -> Something Else

// UPDATE
//  -> Ok
//  -> Invalid Request
//  -> Not Found
//  -> Unauthorized
//  -> Something Else

// DELETE
//  -> Ok
//  -> Not Found
//  -> Unauthorized
//  -> Something Else

type Error struct {
	Status  int       `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Kind    ErrorKind `json:"-"`
}

func (e *Error) Error() string {
	return e.Kind.String() + ": " + e.Message
}

func Conflict(message string) *Error {
	return &Error{
		Status:  http.StatusConflict,
		Message: message,
		Kind:    CONFLICT,
	}
}

func InvalidRequest(message string) *Error {
	return &Error{
		Status:  http.StatusBadRequest,
		Message: message,
		Kind:    INVALID_REQUEST,
	}
}

func NotFound(message string) *Error {
	return &Error{
		Status:  http.StatusNotFound,
		Message: message,
		Kind:    NOT_FOUND,
	}
}

func Unauthorized(message string) *Error {
	return &Error{
		Status:  http.StatusUnauthorized,
		Message: message,
		Kind:    UNAUTHORIZED,
	}
}

func Unknown(err error) *Error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
		Kind:    UNKNOWN,
	}
}

type ErrorKind uint

const (
	CONFLICT ErrorKind = iota
	INVALID_REQUEST
	NOT_FOUND
	UNAUTHORIZED
	UNKNOWN
)

func (k ErrorKind) String() string {
	var prefix string
	switch k {
	case CONFLICT:
		prefix = "conflict"
	case INVALID_REQUEST:
		prefix = "invalid request"
	case NOT_FOUND:
		prefix = "not found"
	case UNAUTHORIZED:
		prefix = "unauthorized"
	case UNKNOWN:
		prefix = "unknown"
	}
	return prefix
}
