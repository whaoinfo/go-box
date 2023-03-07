package gehttp

import (
	"net/http"
	"time"
)

type HTTPServer struct {
	handler *Handler
	svr     *http.Server
}

func (t *HTTPServer) Init(addr string, readTimeout, writeTimeout time.Duration) error {
	t.handler = &Handler{}
	if err := t.handler.init(); err != nil {
		return err
	}

	t.svr = &http.Server{Addr: addr,
		Handler:      t.handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return nil
}

func (t *HTTPServer) Start() error {
	return t.svr.ListenAndServe()
}

func (t *HTTPServer) Stop() error {
	return nil
}

func (t *HTTPServer) AddRoute(path string, handleFunc RouteFunc, methods []string) error {
	t.handler.addRoute(path, handleFunc, methods)
	return nil
}

func (t *HTTPServer) SetWrapRouteExecFunc(f WrapRouteExecFunc) {
	if t.handler == nil {
		return
	}
	t.handler.wrapRouteExecFunc = f
}
