package goway

import (
	"fmt"
//	"log"
	"net/http"
	"reflect"
)

type Goway struct {
	Injector
	handlers []Handler
	action   Handler
	Config   Config
	Logger   Logger
}
type ClassicGoway struct {
	*Goway
	Router
}

// To accept the custom function
// Example:
//  goway.Get("/",function(){
//     //this function is Handler
//  })
// The func() is Handler
type Handler interface{}

func (g *Goway) Handers(handlers ...Handler) {
	g.handlers = make([]Handler, 0)
	for _, handler := range handlers {
		g.Use(handler)
	}
}

func (g *Goway) isHander(handler Handler) error {
	t := reflect.TypeOf(handler)
	if t.Kind() != reflect.Func {
		//params is wrong!
		fmt.Errorf("params is wrong!It params must be an Handler type!")
	}
	return nil
}

func (g *Goway) Use(handler Handler) {
	if g.isHander(handler) == nil {
		g.handlers = append(g.handlers, handler)
	}
}

func (g *Goway) Action(handler Handler) {
	if g.isHander(handler) == nil {
		g.action = handler
	}
}

// Type conversion (*Context)(nil) and (*http.ResponseWriter)(nil):
//      As (* interface{})(nil) Will be nil to the pointer to the interface type
// function NewResponseWriter:
//      It is to implement the HTTP package ResponseWriter interface
func (g *Goway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &context{NewInjector(), g.handlers, g.action, NewResponseWriter(w), 0}
	c.SetParent(g)
	c.MapTo(c, (*Context)(nil))
	c.MapTo(c.rw, (*http.ResponseWriter)(nil))
	c.Map(r)
	c.bootstrap()
}

func (g *Goway) RunOnAddr(addr string) {
	logger := g.Injector.Get(reflect.TypeOf(g.Logger)).Interface().(*Logs)
	logger.Notice(fmt.Sprintf("listening on %s \n", addr))
	logger.logger.Fatalln(http.ListenAndServe(addr, g))
}

func (g *Goway) Run(){
	g.RunOnAddr(g.Config.Get("httpServer").(string)+":"+g.Config.Get("serverPort").(string))
}

func Bootstrap() *ClassicGoway {
	c := InitConfig()
	g := &Goway{Injector: NewInjector(), action: func() {}, Config: c, Logger: InitLogger()}
	r := NewRouter()
	g.Logger.Setloglevel(c.Get("logger").(string))
	g.Logger.Use(c.Logs())
	g.Logger.Print()
	g.Map(g.Logger)
	g.Map(c)
	g.Use(g.Logger.StartLogger())
	g.Use(Recovery(c))
	g.MapTo(r, (*Router)(nil))
	g.Map(defaultReturnHandler()) //The default return by defaultReturnHandler
	g.Action(r.Handler)
	return &ClassicGoway{g, r}

}
