package middleware

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Logger struct {
	l hclog.Logger
}

func NewLogger(l hclog.Logger) *Logger {
	return &Logger{l}
}

func (l *Logger) LogIncomingReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		l.l.Debug("<---- New Reuqeust", "Method", r.Method, "URL:", r.URL.Path, "Remote address", r.RemoteAddr)
		next.ServeHTTP(rw, r)
	})
}
