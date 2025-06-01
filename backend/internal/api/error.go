package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type APIError struct {
	Kind    string `json:"kind"`
	Message any    `json:"message"`
}

func (a APIError) Error() string {
	return fmt.Sprintf("kind: '%s': %s", a.Kind, a.Message)
}

func (a APIError) JSON(w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(a)
}

func ErrResp(w http.ResponseWriter, r *http.Request, code int, msg any) {
	// is error
	if err, ok := msg.(error); ok {
		var apiErr APIError
		if errors.As(err, &apiErr) {
			apiErr.JSON(w)
			return
		}
	}
}
