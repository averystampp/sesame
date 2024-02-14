package sesame

import (
	"log"
	"net/http"
)

type Router struct {
	server *http.Server
	mux    *http.ServeMux
	TLS
}

type TLS struct {
	CertFile string
	KeyFile  string
}

func NewRouter() *Router {
	return &Router{
		mux:    http.NewServeMux(),
		server: &http.Server{},
	}
}

func (r *Router) Get(path string, h Handler) {
	r.mux.Handle("GET "+path, h)
}

func (r *Router) Post(path string, h Handler) {
	r.mux.Handle("POST "+path, h)
}

func (r *Router) Put(path string, h Handler) {
	r.mux.Handle("PUT "+path, h)
}

func (r *Router) Patch(path string, h Handler) {
	r.mux.Handle("PATCH "+path, h)
}

func (r *Router) Options(path string, h Handler) {
	r.mux.Handle("OPTIONS "+path, h)
}

func (r *Router) StartServer(addr string) {
	log.Fatal(http.ListenAndServe(addr, r.mux))
}

func (r *Router) StartServerTLS(addr string) {
	log.Fatal(http.ListenAndServeTLS(addr, r.CertFile, r.KeyFile, r.mux))
}
