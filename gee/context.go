package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Path   string
	Method string

	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (this *Context) Status(code int) {
	this.StatusCode = code
	this.Writer.WriteHeader(code)
}

func (this *Context) SetHeader(key, value string) {
	this.Writer.Header().Set(key, value)
}

func (this *Context) Query(key string) string {
	return this.Req.URL.Query().Get(key)
}

func (this *Context) PostForm(key string) string {
	return this.Req.FormValue(key)
}

func (this *Context) JSON(code int, obj interface{}) {
	this.Status(code)
	this.SetHeader("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(this.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(this.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (this *Context) HTML(code int, html string) {
	this.Status(code)
	this.SetHeader("Content-Type", "text/html; charset=utf-8")
	this.Writer.Write([]byte(html))
}

func (this *Context) String(code int, format string, values ...interface{}) {
	this.Status(code)
	this.SetHeader("Content-Type", "text/plain; charset=utf-8")
	this.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (this *Context) Data(code int, data []byte) {
	this.Status(code)
	this.Writer.Write(data)
}
