package router

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

type RouterInstance struct {
	router *mux.Router
}

func NewRouterInstance() *RouterInstance {
	return &RouterInstance{mux.NewRouter().StrictSlash(true)}
}

func (a *RouterInstance) RegisterHandler(Path string, Handler func(w http.ResponseWriter, r *http.Request), method string) {
	a.router.HandleFunc(Path, Handler).Methods(method)
}

func (a *RouterInstance) Start() {
	http.Handle("/", a.router)
}
