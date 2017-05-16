package mgin

import (
	"fmt"
	"io"
	"net/http"
)

// Engine ...
type Engine struct {
	name string
}

// Context ...
type Context struct {
	res string
	req string
}

// FuncHanlder ...
type FuncHanlder func(context *Context)

// Default ...
func Default() (engine Engine) {

	engine = Engine{
		"mygine",
	}
	return
}

// Get ...
func (engine *Engine) Get(path string, handler func(context *Context)) {

}

// ServeHTTP ...
func (engine *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("get request")
	fmt.Println(req.URL.Query())
	fmt.Println(req.URL.String())
	io.WriteString(res, "hello world")
}

// Run ...
func (engine *Engine) Run(port string) {
	fmt.Println("server starting")
	err := http.ListenAndServe(":9999", engine)
	if err != nil {
		fmt.Println(err)
	}
}

// End ...
func (context *Context) End(body string) {

}
