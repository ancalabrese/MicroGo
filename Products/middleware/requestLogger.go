package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type Logger struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) *Logger {
	return &Logger{l}
}

func (l *Logger) LogIncomingReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("<---[%s] New request received: [%s] --- [%s]", r.Method, r.URL.Path, r.RemoteAddr)
		l.l.Println(msg)
		next.ServeHTTP(rw, r)
	})
}
