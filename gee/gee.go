package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func getKey(method, path string) string {
	return method + "-" + path
}

func (this *Engine) addRoute(method, path string, handler HandleFunc) {
	key := method + "-" + path
	log.Printf("Route %4s - %s", method, path)
	this.router[key] = handler
}

func (this *Engine) GET(pattern string, handler HandleFunc) {
	this.addRoute("GET", pattern, handler)
}

func (this *Engine) POST(pattern string, handler HandleFunc) {
	this.addRoute("POST", pattern, handler)
}

func (this *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := getKey(req.Method, req.URL.Path)
	if handler, ok := this.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (this *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, this)
}
