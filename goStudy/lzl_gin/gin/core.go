package gin

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(ctx Ctx)

type Engine struct {
	router router
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var s strings.Builder
	s.WriteString(req.Method)
	s.WriteString("-")
	s.WriteString(req.URL.Path)
	if handleFunc, ok := e.router.handlers[s.String()]; ok {
		handleFunc(NewCtx(writer, req))
	} else {
		fmt.Fprintf(writer, "404 Not Found: %s\n", req.URL)
	}
}

func New() Engine {
	return Engine{
		router: router{},
	}
}

func (e *Engine) GET(path string, handlerFunc HandlerFunc) {
	e.router.addRoute("GET", path, handlerFunc)
}

func (e *Engine) POST(path string, handlerFunc HandlerFunc) {
	e.router.addRoute("POST", path, handlerFunc)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
