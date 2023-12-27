package httputils

import (
	"net/http"

	"github.com/ecumenos/go-toolkit/random"
)

func ExtractRequestID(r *http.Request) string {
	if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
		return reqID
	}

	return random.GenUUIDString()
}
