package httputils

import (
	"context"
	"net/http"
	"time"

	"github.com/ecumenos/go-toolkit/contextutils"
	"github.com/ecumenos/go-toolkit/primitives"
	"github.com/ecumenos/go-toolkit/random"
)

func ExtractRequestID(r *http.Request) string {
	if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
		return reqID
	}

	return random.GenUUIDString()
}

func GetRequestDuration(ctx context.Context) (int, error) {
	str := contextutils.GetValueFromContext(ctx, contextutils.StartRequestTimestampKey)
	if str == "" {
		return 0, nil
	}
	start, err := primitives.StringToInt64(str)
	if err != nil {
		return 0, err
	}
	diff := time.Now().UnixNano() - start

	return int(diff), nil
}
