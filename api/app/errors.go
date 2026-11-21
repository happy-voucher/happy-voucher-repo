package app

import (
	"net/http"

	"github.com/Piechutowski/volt"
)

type apiError struct {
	Error apiErrorBody `json:"error"`
}

type apiErrorBody struct {
	Code      string        `json:"code"`
	Message   string        `json:"message"`
	RequestID string        `json:"request_id"`
	Details   []volt.Detail `json:"details,omitempty"`
}

func Err(r volt.Request, status int, code, message string) volt.Response {
	return volt.RenderStatus(status, apiError{Error: apiErrorBody{
		Code:      code,
		Message:   message,
		RequestID: r.Header().Get("X-Request-ID"),
	}})
}

func ApiErrors(r volt.Request, err error) volt.Response {
	p, ok := volt.ProblemOf(err)
	if !ok {
		p = volt.Problem{Status: http.StatusInternalServerError, Message: "błąd serwera"}
	}
	return volt.RenderStatus(p.Status, apiError{Error: apiErrorBody{
		Code:      codeForStatus(p.Status),
		Message:   p.Message,
		RequestID: r.Header().Get("X-Request-ID"),
		Details:   p.Details,
	}})
}

func codeForStatus(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "BAD_REQUEST"
	case http.StatusUnauthorized:
		return "UNAUTHENTICATED"
	case http.StatusForbidden:
		return "FORBIDDEN"
	case http.StatusNotFound:
		return "NOT_FOUND"
	case http.StatusConflict:
		return "CONFLICT"
	case http.StatusUnprocessableEntity:
		return "VALIDATION_FAILED"
	}
	return "INTERNAL"
}
