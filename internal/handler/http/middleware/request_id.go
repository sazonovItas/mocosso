package httpmiddleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"

type ctxRequestIdKey struct{}

func RequestID() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(RequestIDHeader)
			if requestId == "" {
				requestId = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), ctxRequestIdKey{}, nextRequestID())
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(ctxRequestIdKey{}).(string)
}

func nextRequestID() string {
	return uuid.NewString()
}
