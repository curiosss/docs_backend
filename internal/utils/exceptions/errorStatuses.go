package exceptions

import (
	"net/http"

	"github.com/goccy/go-json"
)

// ResponseError defines a detailed error structure
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// ResponseError returns the error as a JSON string
func (e *ResponseError) Error() string {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return `{"code":500,"message":"Internal Server ResponseError","details":"Failed to encode error to JSON"}`
	}
	return string(jsonData)
}

// Predefined custom errors for common HTTP status codes
var (
	ErrBadRequest = &ResponseError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Details: "Maglumatlaryň dogrulygyny barlaň!",
	}
	ErrUnauthorized = &ResponseError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Details: "Ulgama giriň!",
	}
	ErrForbidden = &ResponseError{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
		Details: "Sizde bu resursa elýeterlik ýok.",
	}
	ErrNotFound = &ResponseError{
		Code:    http.StatusNotFound,
		Message: "Not Found",
		Details: "Tapylmady",
	}
	ErrConflict = &ResponseError{
		Code:    http.StatusConflict,
		Message: "Conflict",
		Details: "The request could not be completed due to a conflict with the current state of the resource.",
	}
	ErrInternalServerError = &ResponseError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server ResponseError",
		Details: "Ýalňyşlyk ýüze çykdy, gaýtadan synanyşyp görüň!",
	}
	ErrUnprocessableEntity = &ResponseError{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
		Details: "The server understands the content type of the request but was unable to process the contained instructions.",
	}
	ErrTooManyRequests = &ResponseError{
		Code:    http.StatusTooManyRequests,
		Message: "Too Many Requests",
		Details: "The user has sent too many requests in a given amount of time.",
	}
)

func NewResponseError(baseError *ResponseError, err error) *ResponseError {
	return &ResponseError{
		Code:    baseError.Code,
		Message: baseError.Message,
		Details: err.Error(),
	}
}
