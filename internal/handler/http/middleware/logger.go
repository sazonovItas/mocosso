package httpmiddleware

import (
	"net/http"

	"github.com/sazonovItas/mocosso/pkg/logger"
	"go.uber.org/zap"
)

func LoggerToContext(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := logger.ToContext(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func Logger(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logWriter := &loggerResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(logWriter, r)
		}

		return http.HandlerFunc(fn)
	}
}

type loggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggerResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}
