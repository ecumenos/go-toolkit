package httputils

import (
	"context"
	"net/http"
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/ecumenos/go-toolkit/contextutils"
	"github.com/ecumenos/go-toolkit/primitives"
	"github.com/ecumenos/go-toolkit/random"
	"github.com/ecumenos/go-toolkit/timeutils"
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

type Metadata struct {
	RequestID string `json:"requestId"`
	Duration  int    `json:"duration"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

func GetMetadata(ctx context.Context) (*Metadata, error) {
	duration, err := GetRequestDuration(ctx)
	if err != nil {
		return nil, err
	}

	return &Metadata{
		RequestID: contextutils.GetValueFromContext(ctx, contextutils.RequestIDKey),
		Timestamp: timeutils.TimeToString(time.Now()),
		Duration:  duration,
		Version:   versioninfo.Short(),
	}, nil
}
