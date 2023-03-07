package gehttp

import "net/http"

type WrapRouteExecFunc func(hdFunc RouteFunc, w http.ResponseWriter, r *http.Request)
type RouteFunc func(w http.ResponseWriter, r *http.Request) error

type RouteInfo struct {
	MethodList []string
	methodSet  map[string]bool
	HandleFunc RouteFunc
}
