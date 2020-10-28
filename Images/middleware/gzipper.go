package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type Gzipper struct {
}

// Gzipper middleware cheks whether gzip is supported by the client
// and if true compresses the image before streaming it out
func (g *Gzipper) GzipperMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//Client doesn't support gzip compression
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(rw, r)
		}
		GzipRW := newGzipResponseWriter(rw)
		GzipRW.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(GzipRW, r)
		defer GzipRW.Flush()
	})
}

type GzipResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func newGzipResponseWriter(rw http.ResponseWriter) *GzipResponseWriter {
	gw := gzip.NewWriter(rw)
	return &GzipResponseWriter{rw: rw, gw: gw}
}

func (gz *GzipResponseWriter) Header() http.Header {
	return gz.rw.Header()
}

func (gz *GzipResponseWriter) Write(d []byte) (int, error) {
	return gz.gw.Write(d)
}

func (gz *GzipResponseWriter) WriteHeader(code int) {
	gz.rw.WriteHeader(code)
}
func (gz *GzipResponseWriter) Flush() {
	gz.gw.Flush()
	gz.gw.Close()
}
