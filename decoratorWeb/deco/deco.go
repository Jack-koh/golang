package deco

import "net/http"

type DecoratorFunc func(w http.ResponseWriter, r *http.Request, handler http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

func (s *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.fn(w, r, s.h)
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{fn, h}
}
