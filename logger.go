package goway

import (
	"log"
	"net/http"
	"time"
)

func Logger() Handler {
	return func(res http.ResponseWriter, req *http.Request, c Context, log *log.Logger) {
		start := time.Now()

		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}

		log.Printf("Started %s %s for %s", req.Method, req.URL.Path, addr)

		rw := res.(ResponseWriter)
		c.Next()

		log.Printf("Completed %v %s, Content-Length: %v bytes in %v\n", rw.Status(), http.StatusText(rw.Status()), rw.Size(), time.Since(start))
	}
}
