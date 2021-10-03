package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	liberrors "github.com/lab1/errors-logs/pkg/lib/errors"
)

func (a *API) handleError(w http.ResponseWriter, err error) {
	var codeErr liberrors.Err
	if !errors.As(err, &codeErr) {
		// Client shouldn't see unexpected errors
		codeErr = liberrors.InternalError
	}

	a.logger.WithField("code", codeErr.Code()).Error(err.Error())
	_ = encodeError(w, codeErr)
}

type errorResponse struct {
	Code   string          `json:"code"`
	Errors []errorResponse `json:"errors,omitempty"`
	Field  string          `json:"field,omitempty"`
}

func encodeError(w http.ResponseWriter, err liberrors.Err) error {
	w.WriteHeader(errorCode(err))

	var resp interface{}
	b := liberrors.NewBundle()
	if errors.As(err, &b) {
		// bundle will render list of errors
		resp = errorResponse{Code: err.Code(), Errors: toErrorResponses(b)}
	} else {
		resp = toErrorResponse(err)
	}

	if encodeErr := json.NewEncoder(w).Encode(resp); encodeErr != nil {
		return fmt.Errorf("encode json: %w", encodeErr)
	}

	return nil
}

func errorCode(err liberrors.Err) int {
	switch err {
	case liberrors.InternalError:
		return http.StatusInternalServerError
	case liberrors.Unauthorized:
		return http.StatusUnauthorized
	case liberrors.PermissionDenied:
		return http.StatusForbidden
	case liberrors.NotFound:
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}

func toErrorResponses(b *liberrors.Bundle) []errorResponse {
	resp := make([]errorResponse, 0, len(b.List()))
	for _, e := range b.List() {
		resp = append(resp, toErrorResponse(e))
	}

	return resp
}

func toErrorResponse(err error) errorResponse {
	var (
		valErr  = &liberrors.ValidationErr{}
		codeErr = &liberrors.Err{}
	)

	switch {
	case errors.As(err, valErr):
		return errorResponse{Code: valErr.Code(), Field: valErr.Field()}
	case errors.As(err, codeErr):
		return errorResponse{Code: codeErr.Code()}
	default:
		// Client shouldn't see unexpected errors
		return errorResponse{Code: liberrors.InternalError.Code()}
	}
}
