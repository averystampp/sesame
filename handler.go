package sesame

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Context struct {
	resp http.ResponseWriter
	req  *http.Request
	db   *sql.DB
}

func (c *Context) Response() http.ResponseWriter {
	return c.resp
}

func (c *Context) Request() *http.Request {
	return c.req
}

func (c *Context) DB() (*sql.DB, error) {
	if c.db == nil {
		return nil, fmt.Errorf("no db initialized")
	}
	return c.db, nil
}

type Handler func(Context) error

// Make Handler a http.Handler interface to satisfy directly passing handlers in to mux
// and for httptest
func (h Handler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	for _, m := range middleWare {
		err := m(Context{resp: wr, req: r, db: DB})
		if err != nil {
			wr.Write([]byte(err.Error()))
			return
		}
	}

	if err := h(Context{resp: wr, req: r, db: DB}); err != nil {
		wr.Write([]byte(err.Error()))
	}
}
