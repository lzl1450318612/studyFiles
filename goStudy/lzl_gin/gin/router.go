package gin

import (
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
}

func (r *router) addRoute(method string, path string, handlerFunc HandlerFunc) {
	var s strings.Builder
	s.WriteString(method)
	s.WriteString("-")
	s.WriteString(path)
	r.handlers[s.String()] = handlerFunc
}
