package httplogger

import (
	"bufio"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/spf13/viper"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

// fix the error: unable to upgrade *main.statusRecorder to websocket websocket: response does not implement http.Hijacker
// FIXME: ok?
func (rec *statusRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rec.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, 200}
		next.ServeHTTP(&rec, r)
		if rec.status/100 != 2 || viper.GetBool("DEBUG") {
			log.Printf("%v %v\n", r.URL, rec.status)
		}
	})
}
