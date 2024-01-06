package httputils

import (
	"context"
	"errors"
	"net/http"
	"strings"
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

func ExtractJWTBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("token prefix is missing")
	}

	return authHeaderParts[1], nil
}
