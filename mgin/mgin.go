package mgin

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

// Run ...
func (engine *Engine) Run(port string) {

}

func (context *Context) End(body string) {

}
