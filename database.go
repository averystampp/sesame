package sesame

import (
	"database/sql"
)

type DBOpts struct {
	User, Password, Host, Dbname string
	Port                         int
}

var DB *sql.DB

func (r *Router) SetDB(f func() *sql.DB) {
	DB = f()
}
