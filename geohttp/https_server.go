package geohttp

import (
	"crypto/tls"
	"net/http"
	"time"
)

type HTTPSServer struct {
	HTTPServer
	certFile, keyFile string
}

func (t *HTTPSServer) Init(addr, certFile, keyFile string, readTimeout, writeTimeout time.Duration) error {
	t.handler = &Handler{}
	if err := t.handler.init(); err != nil {
		return err
	}

	t.certFile = certFile
	t.keyFile = keyFile

	t.svr = &http.Server{
		Addr:         addr,
		Handler:      t.handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		TLSConfig:    &tls.Config{},
	}

	return nil
}

func (t *HTTPSServer) Start() error {
	return t.svr.ListenAndServeTLS(t.certFile, t.keyFile)
}
