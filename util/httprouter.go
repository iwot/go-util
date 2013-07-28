package util

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
)

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{
		pathMatch:  NewPathMatch(),
		basePath:   "",
		actions:    make(map[string]ActionFunc),
		errorFuncs: make(map[string]ErrorFunc),
	}
}

type ActionFunc func(map[string]string, http.ResponseWriter, *http.Request) (string, ActionError)

func NewActionError(code string, message string) ActionError {
	return ActionError{code, message}
}

type ActionError struct {
	code    string
	message string
}

func (e ActionError) Error() string {
	return e.message
}
func (e ActionError) Code() string {
	return e.code
}

func IsActionError(err ActionError) bool {
	if err.code == "" && err.message == "" {
		return false
	}
	return true
}

type ErrorFunc func(string, http.ResponseWriter, *http.Request)
type Executor struct {
	params map[string]string
	action ActionFunc
}

func (e *Executor) Exec(w http.ResponseWriter, r *http.Request) (string, ActionError) {
	return e.action(e.params, w, r)
}

type HttpRouter struct {
	pathMatch  *PathMatch
	basePath   string
	actions    map[string]ActionFunc
	errorFuncs map[string]ErrorFunc
}

func (self *HttpRouter) SetBasePath(basePath string) *HttpRouter {
	self.basePath = basePath
	return self
}

func (self *HttpRouter) SetError(errorCode string, errorFunc ErrorFunc) *HttpRouter {
	self.errorFuncs[errorCode] = errorFunc
	return self
}

func (self *HttpRouter) SetRoute(pathPattern string, action ActionFunc) *HttpRouter {
	defaults := make(map[string]string)
	defaults["action"] = "[a-z0-9_]+"
	err := self.pathMatch.Parse(pathPattern, defaults)
	if err != nil {
		panic(err)
	}
	self.actions[pathPattern] = action
	return self
}

func (self *HttpRouter) removeBasePathFrom(path string) string {
	if len(path) >= len(self.basePath) && path[0:len(self.basePath)] == self.basePath {
		p := path[len(self.basePath):]
		return p
	} else {
		return ""
	}
}

func (self *HttpRouter) GetAction(urlPath string) (Executor, error) {
	path := self.removeBasePathFrom(urlPath)
	pathPattern, matches, ok := self.pathMatch.Match(path)
	e := Executor{}
	if ok {
		e.params = matches
		e.action = self.actions[pathPattern]
		return e, nil
		/*result := func(w http.ResponseWriter, r *http.Request) (string, error) {
			return self.actions[pathPattern](matches, w, r)
		}
		return result*/
		//return self.actions[pathPattern](self, r)
	}
	return e, errors.New("action not found")
}

func (self *HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := self.GetAction(r.URL.Path)
	if err == nil {
		message, err := f.Exec(w, r)
		if IsActionError(err) {
			self.errorFuncs[err.Code()](err.Error(), w, r)
		}
		log.Printf("result[%s]", message)
	} else {
		defaultFunc, ok := self.errorFuncs["default"]
		if ok {
			defaultFunc(r.URL.Path, w, r)
		} else {
			fmt.Fprintf(w, "Error, %q\n", html.EscapeString(r.URL.Path))
		}
	}
	/*path := self.removeBasePathFrom(r.URL.Path)
	pathPattern, matches, ok := self.pathMatch.Match(path)
	fmt.Println("path = ", path)
	fmt.Println("pathPattern = ", pathPattern)
	if ok {
		self.actions[pathPattern](self, r)
	}*/
}
