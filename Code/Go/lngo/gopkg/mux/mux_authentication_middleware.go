package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gorilla/mux"
)

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			// http.Error(w, "Forbidden", http.StatusForbidden)

			// for debug and pass; temp add
			next.ServeHTTP(w, r)
		}
	})
}

func hi(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(res, "hi")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Do something here
	})
	amw := authenticationMiddleware{make(map[string]string)}
	amw.Populate()
	r.Use(amw.Middleware)

	r.HandleFunc("/hi", hi)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 31 * time.Second,
		ReadTimeout:  31 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
