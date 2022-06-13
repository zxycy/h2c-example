package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"time"
)

var PemPath = "./TLS/server.crt"
var KeyPath = "./TLS/server.key"

func NewServer(bindAddr string, handler http.Handler) (server *http.Server, err error) {
	if handler == nil {
		return nil, errors.New("server needs handler to handle request")
	}

	h2Server := &http2.Server{
		// TODO: extends the idle time after re-use openapi client
		IdleTimeout: 1 * time.Millisecond,
	}

	server = &http.Server{
		Addr:    bindAddr,
		Handler: h2c.NewHandler(handler, h2Server),
	}

	return
}

func StartServer(scheme string) {
	route := gin.New()
	AddService(route)
	server, _ := NewServer("127.0.0.1:8888", route)
	if scheme == "http" {
		server.ListenAndServe()
	} else if scheme == "https" {
		server.ListenAndServeTLS(PemPath, KeyPath)
	}
}
