package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type RouteDefinitionList map[string]map[string]http.HandlerFunc
type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

func AppendListToMux(mux *http.ServeMux, list *RouteDefinitionList, mwFunc MiddlewareFunc) {

	if list == nil {
		panic("expected given list to not be nil")
	}

	for method, routes := range *list {
		for route, handler := range routes {
			if mwFunc == nil {
				mux.HandleFunc(fmt.Sprintf("%s %s", method, route), handler)

			} else {
				mux.HandleFunc(fmt.Sprintf("%s %s", method, route), mwFunc(handler))

			}
		}
	}
}

func StatCatResponder(contentType string, filenames ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", contentType)
		w.WriteHeader(200)
		for _, file := range filenames {
			if f, err := os.OpenFile(file, os.O_RDONLY, 0); err == nil {
				io.Copy(w, f)
				f.Close()
			} else {
				RespondWithError(w, 500, err)

			}
		}
	}
}
