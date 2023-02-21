package geohttp

import (
	"fmt"
	"io"
	"net/http"
)

type Handler struct {
	routeInfoMap      map[string]*RouteInfo
	wrapRouteExecFunc WrapRouteExecFunc
}

func (t *Handler) init() error {
	t.routeInfoMap = make(map[string]*RouteInfo)
	return nil
}

func (t *Handler) addRoute(path string, handleFunc RouteFunc, methods []string) {
	routeInfo := &RouteInfo{
		methodSet:  make(map[string]bool),
		HandleFunc: handleFunc,
	}
	for _, method := range methods {
		routeInfo.methodSet[method] = true
	}

	t.routeInfoMap[path] = routeInfo
}

func (t *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	routeInfo, exist := t.routeInfoMap[path]
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if ok := routeInfo.methodSet[r.Method]; !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if t.wrapRouteExecFunc != nil {
		t.wrapRouteExecFunc(routeInfo.HandleFunc, w, r)
	} else {
		routeInfo.HandleFunc(w, r)
	}

}

func BuildDefaultJsonResponse(w http.ResponseWriter, r *http.Request, code int) error {
	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, fmt.Sprintf(`{"code": %d}`, code))
	return err
}
