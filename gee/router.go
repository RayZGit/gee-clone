package gee

type Router struct {
	handler map[string]HandlerFunc
}

func getKey(method, path string) string {
	return method + "-" + path
}

func NewRouter() *Router {
	return &Router{handler: make(map[string]HandlerFunc)}
}

func (this *Router) add(method, pattern string, handler HandlerFunc) {
	this.handler[getKey(method, pattern)] = handler
}

func (this *Router) handle(c *Context) {
	key := getKey(c.Method, c.Path)
	if handler, ok := this.handler[key]; ok {
		handler(c)
	}
}
