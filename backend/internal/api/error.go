package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	KindInternalServerError = "INTERNAL_SERVER_ERROR"
)

type APIError struct {
	Kind   string `json:"kind"`
	Detail any    `json:"detail"`
}

func NewAPIError(kind string, detail any) APIError {
	return APIError{Kind: kind, Detail: detail}
}

func (a APIError) Error() string {
	return fmt.Sprintf("kind: '%s': %s", a.Kind, a.Detail)
}

func (a APIError) JSON(w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(a)
}

func ErrResp(w http.ResponseWriter, r *http.Request, code int, msg any) {
	w.WriteHeader(code)

	switch x := msg.(type) {
	case APIError:
		x.JSON(w)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		APIError{Kind: KindInternalServerError, Detail: msg}.JSON(w)
		return
	}

}
