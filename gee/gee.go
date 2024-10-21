package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (this *Engine) addRoute(method, path string, handler HandlerFunc) {
	log.Printf("Register Route %4s - %s", method, path)
	this.router.add(method, path, handler)
}

func (this *Engine) GET(pattern string, handler HandlerFunc) {
	this.addRoute("GET", pattern, handler)
}

func (this *Engine) POST(pattern string, handler HandlerFunc) {
	this.addRoute("POST", pattern, handler)
}

func (this *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//this.router.handle(w, req)
	c := NewContext(w, req)
	this.router.handle(c)

}

func (this *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, this)
}
